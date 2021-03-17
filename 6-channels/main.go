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

	//for {
	//	go checkLink(<-c, c)
	//}

	//for l := range c {
	//	go checkLink(l, c)
	//}

	// With Function literal to add sleep
	// Make sure the pause is not in the main Routine
	// and not incorrectly in the checkLink function
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			go checkLink(link, c)
		}(l)
	}

	for {
		checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {
	_, error := http.Get(link)

	if error != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}
