package monitor

import (
	"log"
	"net/http"
	"sync"
	"time"
	"doga/alert"
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
				alert.SendMattermostAlert(duration)
			}
		}
	}()
}
