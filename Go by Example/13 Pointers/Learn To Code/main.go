package main

/*
When To Use Pointers in Programming - Go Lang Pointers by Learn To Code
https://www.youtube.com/watch?v=1YEufRd3_Fk
*/

import "fmt"

// func foo(y int) {
// 	fmt.Println("first print in foo func", y)
// 	y = 43
// 	fmt.Println("second print in foo func", y)
// }

// func main() {
// 	x := 0
// 	foo(x)
// 	fmt.Println("print in main func", x)

// }

func foo(y *int) {
	fmt.Println("func foo y before: ", y, "mem place")
	fmt.Println("func foo y before: ", *y, "dereferenced (value)")
	*y = 43
	fmt.Println("func foo y after: ", y, "mem place")
	fmt.Println("func foo y after: ", *y, "dereferenced (value)")
}

func main() {
	x := 0
	fmt.Println("func main x before: ", &x, "mem place")
	fmt.Println("func main x before: ", x, "dereferenced (value)")
	foo(&x)
	fmt.Println("func main x after: ", &x, "mem place")
	fmt.Println("func main x after: ", x, "dereferenced (value)")

}
