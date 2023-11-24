package monitor

import (
	"doga/alert"
	"io/ioutil"
	"log"
	"net/http"
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

func init() {
	go func() {
		for {
			time.Sleep(10 * time.Second)
			mu.Lock()
			duration := time.Since(lastRequestTime)
			mu.Unlock()
			if duration > 1*time.Minute {
				log.Printf("No requests in the last %v", duration)
				ms := alert.InitializeMattermostSender("jnhq1erhftyz9dt7y7qzxcg38r")
				ms.SendAlert("No requests in the last 1 minute.")
			}
		}
	}()
}
