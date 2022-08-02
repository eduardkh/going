package main

import (
	"fmt"
)

func fizzbuzzpagaz(max int) {
	for i := 1; i <= 100; i++ {

		// fmt.Printf("> i is :%v\n", i) // for debug
		output := ""
		if i%3 == 0 {
			output += "Fizz"
		}
		if i%5 == 0 {
			output += "Buzz"
		}
		if i%7 == 0 {
			output += "Pagaz"
		}
		if output != "" {
			fmt.Println(output)
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fizzbuzzpagaz(100)
}
