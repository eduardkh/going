package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed standards-oui.ieee.org.txt
var ouiList string

func main() {
	// split the string to lines
	lines := strings.Split(string(ouiList), "\n")

	// remove the last item from the lines slice
	// which is empty
	lines = lines[:len(lines)-1]
	// get the line that contains 4838B6
	for _, line := range lines {
		if strings.Contains(line, "4838B6") {
			// split the line to individual elements
			words := strings.Fields(line)
			fmt.Println(words[0], ": ", strings.Join(words[3:], " "))
		}
	}
}
