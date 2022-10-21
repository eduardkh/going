package main

import "fmt"

func main() {
	hi := "hi"
	fmt.Println(hi)
	// "&" point to memory location and "*" change the value
	*&hi = "bye"
	fmt.Println(hi)

	// "&" point to memory location and "*" change the value in two steps
	aa := "aa"
	fmt.Println(aa)
	bb := &aa
	*bb = "bb"
	fmt.Println(aa)
}
