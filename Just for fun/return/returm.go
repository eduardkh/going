package main

import "fmt"

func main() {

	fmt.Println(myFunc())

}

func myFunc() string {
	a := "Just a string"
	b := 5
	return fmt.Sprintf("String => %v :: Int => %v", a, b)
}
