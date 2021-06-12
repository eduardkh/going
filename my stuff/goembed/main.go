package main

import (
	_ "embed"
	"fmt"
)

//go:embed sample.txt
var text string

func main() {
	fmt.Print(text)
}
