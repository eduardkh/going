package main

import "fmt"

func main() {
	nameNum := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	// num, exists := nameNum["one"]
	// fmt.Println(num, exists)

	if v, exists := nameNum["two"]; exists {
		fmt.Printf("number %v exists\n", v)
	} else {
		fmt.Println("number doesn't exist")
	}

	// num5, exists5 := nameNum["five"]
	// fmt.Println(num5, exists5)
	if v5, exists5 := nameNum["five"]; exists5 {
		fmt.Printf("number %v exists\n", v5)
	} else {
		fmt.Println("number doesn't exist")
	}
}
