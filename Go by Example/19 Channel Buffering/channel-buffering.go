package main

import (
	"fmt"
)

// https://gobyexample.com/channel-buffering

func main() {

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
