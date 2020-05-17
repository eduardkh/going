package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
	suka := "suka"
	fmt.Println("suka initial value:", suka)
	fmt.Println("suka memory address:", &suka)
	fmt.Println("suka value in memory address:", *&suka)

    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i)
}