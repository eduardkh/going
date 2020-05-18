package main

import (
	"fmt"
	"time"
)

// https://gobyexample.com/goroutines
// feels like pythons multithreading but baked-in

func f(from string) {
	for i := 1; i <= 10; i++ {
		// time.Sleep(time.Second) // if sleeping can't see goroutines
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
