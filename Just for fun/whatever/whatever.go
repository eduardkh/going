package main

import "fmt"

func funcToVar() string {
	return "func To Var"
}

func main() {
	t := true
	st := fmt.Sprintf("type of t is %T", t)
	fmt.Println(st)

	fts := funcToVar()
	fmt.Println(fts)

}
