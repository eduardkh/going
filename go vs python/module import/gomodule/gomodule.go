// https://www.digitalocean.com/community/tutorials/how-to-use-go-modules
package main

import (
	"fmt"
	"mymodule/gopackage"
)

func main() {
	fmt.Println("Hello, world from Modules!")
	fmt.Println(gopackage.Message)
}
