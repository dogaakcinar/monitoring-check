package monitor

import (
	"doga/alert"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	mu sync.Mutex
	// Start with the last request time as now
	lastRequestTime = time.Now()
)

func Handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	lastRequestTime = time.Now()

	// Log the request method and URL path
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read request body: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Printf("Body: %s", body)
}

func StartMonitor() {
	alertIntervalStr := os.Getenv("ALERT_INTERVAL_SECOND")

	if alertIntervalStr == "" {
		log.Fatalf("ALERT_INTERVAL_SECOND is not set or is an empty string")
	}

	var alertInterval int
	var err error
	if alertInterval, err = strconv.Atoi(alertIntervalStr); err != nil {
		log.Fatalf("Failed to parse ALERT_INTERVAL_SECOND: %v", err)
	}

	go func() {
		for {
			time.Sleep(10 * time.Second)
			mu.Lock()
			duration := time.Since(lastRequestTime)
			mu.Unlock()
			if duration > time.Duration(alertInterval)*time.Second {
				log.Printf("No requests in the last %v", duration)
				ms := alert.InitializeMattermostSender(os.Getenv("MATTERMOST_HOOK_ID"), os.Getenv("MATTERMOST_SERVER_URL"))
				ms.SendAlert("No requests in the last 1 minute.")
			}
		}
	}()
}
