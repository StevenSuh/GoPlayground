package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for link := range c {
		// go checkLink(link, c, true)
		go func(link string) {
			time.Sleep(time.Duration(5) * time.Second)
			checkLink(link, c)
		}(link)
	}
}

func checkLink(link string, c chan string) {
	start := time.Now()

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is DOWN")
		// c <- "DOWN"
		c <- link
		return
	}

	end := time.Now()
	fmt.Println(link, "is OK - ", end.Sub(start))
	// c <- "OK"
	c <- link
}
