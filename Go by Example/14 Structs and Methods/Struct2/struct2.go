package main

import "fmt"

var config = struct {
	name string
	age  int
}{
	name: "Bob",
	age:  20,
}

func main() {
	fmt.Print(config)
}
