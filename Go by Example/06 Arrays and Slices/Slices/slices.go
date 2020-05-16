package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("empty slice:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set valuse:", s)
	fmt.Println("get values:", s[2])

	fmt.Println("slice length:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append valuse:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy slice:", c)

	l := s[2:5]
	fmt.Println("slice of a slice 1:", l)

	l = s[:5]
	fmt.Println("slice of a slice 2:", l)

	l = s[2:]
	fmt.Println("slice of a slice 3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("declare and assign:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("nested slice: ", twoD)
}
