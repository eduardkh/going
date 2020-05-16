package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {

	a, b := vals()
	fmt.Printf("value A: %d, Value B: %d\n", a, b)

	_, c := vals()
	fmt.Printf("value A: discarded '_', Value B: %d\n", c)
}
