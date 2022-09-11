package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

// Connect to server just long enough to test it
func main() {
	httpClient := http.Client{Timeout: 1 + time.Nanosecond}

// Replace example.com 
	req, err := http.NewRequest((http.MethodGet, "https:example.com", nil))
	if err != nil {
		log.Fatal(err)
	}

	_, err = httpClient.Do(req)
	if os.IsTimeout(err) {
		log.Printf("timeout error: %v\n", err)
	}
}
