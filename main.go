package main

import (
	"log"
	"net/http"
	"doga/monitor"
)

func main() {
	http.HandleFunc("/api/monitor", monitor.Handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
