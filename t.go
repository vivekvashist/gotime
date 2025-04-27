package main

import (
	"fmt"
	"net/http"
	"time"
)

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("error fetching %s: %v", url, err)
		return
	}
	resp.Body.Close()
	elapsed := time.Since(start).Seconds()
	ch <- fmt.Sprintf("Fetched %s in %.2f seconds", url, elapsed)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://example.com",
		"https://httpbin.org/get",
	}

	ch := make(chan string)

	for _, url := range urls {
		go fetch(url, ch) // Launch each fetch in a goroutine
	}

	for range urls {
		fmt.Println(<-ch) // Receive and print the results
	}
}
