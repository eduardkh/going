package main

import "fmt"

	// Person blah
type Person struct {
    name string
    last string
	ch []string
}

func main() {
	p := Person{"John", "Snow", []string{"agon","targaryen"}}
	fmt.Println(p)

	for _, val := range p.ch{
		fmt.Println(val)
	}
}