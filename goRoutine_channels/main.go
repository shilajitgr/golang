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
		"http://gmail.com",
		"http://golang.org",
		"http://amazon.in",
	}

	c := make(chan string)
	// c := make(chan interface{})	:	can accept value of any type
	for _, link := range links {
		go checklink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// for { // infinite loop, this will be moving onto the next iteration only when channel has some data
	for l := range c { // this here allows to flow to enter the body only when an item has been sent in the channel
		// the value in the channel gets consumed by l

		// fmt.Println(<-c) // blocking call, waits until a msg is put in the channel
		// go checklink(l, c)
		go func(link string) {
			time.Sleep(5 * time.Second) // pauses the flow for 5 secs
			checklink(link, c)
		}(l) // paranthesis perform the same function as those in a func call, eg: checklink(...)
	}
}

func checklink(site string, c chan string) {
	_, err := http.Get(site)

	if err != nil {
		fmt.Println(site, "might be down")
		//c <- "might be down"
		c <- site
	} else {
		fmt.Println(site, "is up")
		// c <- "its up"
		c <- site
	}
}
