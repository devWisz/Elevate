package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://sarjakkhanal.com",
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
		"http://this-site-should-fail.com",
	}

	c := make(chan string)

	for _, link := range links {
	
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	_, err := client.Get(link)
	if err != nil {
		c <- "❌ " + link + " might be down! Error: " + err.Error()
		return
	}

	c <- "✅ " + link + " is up and running!"
}