package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request Failed")

func main() {
	// var results = map[string]string{}
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	// results["hello"] = "Hello"

	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		// fmt.Println(<-c)
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, ": ", status)
	}
}

func hitURL(url string, c chan<- requestResult) {
	// fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		// fmt.Println(err, resp.StatusCode)
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}
