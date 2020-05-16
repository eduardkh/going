package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("empty array:", a)

	a[4] = 100
	fmt.Println("set value:", a)
	fmt.Println("get value:", a[4])

	fmt.Println("length of the array:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("declare and assign:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("nested array: ", twoD)
}
