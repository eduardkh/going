package main

import "fmt"

// https://gobyexample.com/channels
// channels allow you to communicate between 2 goroutines

func main() {

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
