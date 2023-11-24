package main

import (
	"doga/monitor"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/monitor", monitor.Handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
