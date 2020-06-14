package main

import "fmt"

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func main() {
	str := "שלום צ'ורני"
	strRev := reverse(str)
	fmt.Println(strRev)
}