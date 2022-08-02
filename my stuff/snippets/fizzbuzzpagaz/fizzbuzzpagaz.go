package main

import "fmt"

func fizzbuzzpagaz(max int) {
	for i := 1; i < max+1; i++ {
		// fmt.Printf("> i is :%v\n", i) // for debug
		switch {
		case (i%15 == 0 && i%5 == 0 && i%3 == 0):
			fmt.Println("FizzBuzzPagaz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		default:
			fmt.Println(i)
		}
	}
}

func main() {
	fizzbuzzpagaz(100)
}
