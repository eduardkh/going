package main

import (
	"fmt"
	"time"
)

// https://gobyexample.com/channel-synchronization

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second * 1)
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	<-done
}
