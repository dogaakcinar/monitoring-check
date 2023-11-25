package main

import (
	"doga/monitor"
	"log"
	"net/http"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	monitor.StartMonitor()

	http.HandleFunc("/api/monitor", monitor.Handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
