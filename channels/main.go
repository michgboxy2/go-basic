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
		go checkLink(link, c) //only use the go keyword in front of function calls
	}

	// fmt.Println(<-c) //receive a message from the channel
	// fmt.Println(<-c) //receive a message from the channel
	// fmt.Println(<-c) //receive a message from the channel
	// fmt.Println(<-c) //receive a message from the channel
	// fmt.Println(<-c) //receive a message from the channel

	// fmt.Println(<-c) //receive a message from the channel

	// for { //infinite loop so the sites can be getting pinged continously
	// 	// fmt.Println(<-c)
	// 	go checkLink(<-c, c)
	// }

	for l := range c { //infinite loop so the sites can be getting pinged continously
		go func(link string) { //function literal
			time.Sleep(5 * time.Second) //wait 5 seconds
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		// c <- "Might be down i think" //send a message to the channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// c <- "Yup it's up" //send a message to the channel
	c <- link
}
