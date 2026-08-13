package main

import (
	"fmt"
	"math"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/time/utc", timeHandler)
	http.HandleFunc("/time/days-until", daysUntilHandler)
	http.HandleFunc("/", rootHandler)

	fmt.Println("HTTP time server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	var currentTime string
	if r.URL.Path == "/time/utc" {
		currentTime = time.Now().UTC().Format(time.RFC1123)
	} else {
		currentTime = time.Now().Format(time.RFC1123)
	}
	w.Write([]byte(currentTime + "\n"))
}

func daysUntilHandler(w http.ResponseWriter, r *http.Request) {
	targetStr := r.URL.Query().Get("date")
	currentTime := time.Now().Format("2006-01-02")
	target, err := time.Parse("2006-01-02", targetStr)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}
	daysUntil := target.Sub(time.Now()).Hours() / 24
	w.Write([]byte(fmt.Sprintf("Starting from %s, there are %d days until %s\n", currentTime, int(math.Floor(daysUntil)), targetStr)))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Use GET /time to get the current time\n"))
}
