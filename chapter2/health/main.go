package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

// HealthCheck API return date time to client
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	io.WriteString(w, currentTime.String())
}

func main() {
	http.HandleFunc("/health", HealthCheck)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
