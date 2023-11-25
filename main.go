package main

import (
	"doga/monitor"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fmt.Println(os.Getenv("ALERT_INTERVAL_SECOND"))

	monitor.StartMonitor()

	http.HandleFunc("/api/monitor", monitor.Handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
