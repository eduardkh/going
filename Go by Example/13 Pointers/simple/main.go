package main

import "fmt"

func main() {
	// initiate record as 15
	record := 15
	// point to it's memory location
	ptr := &record
	// print initial value
	fmt.Println(record)
	// print value memory location
	fmt.Println(ptr)
	// print value in memory location
	fmt.Println(*ptr)
	// dereferencing memory location and assigning a new value
	*ptr = 17
	// print initial record value (changed)
	fmt.Println(record)

}
