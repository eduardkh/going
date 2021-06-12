package main

import (
	"fmt"
	"math"
)

/*
same request format/endpoint/method
one interface
two different shapes

for an interface to work we need:
1. struct(s) - here we have 2 shapes
2. method - they share the same name but return different answers
3. interface(s) - the methods must have the same name to work with the interface
4. calling function - must call the interface (the only dynamic type)


*/

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("======================")
	fmt.Println("g geometry: ", g)
	fmt.Println("g.area: ", g.area())
	fmt.Println("g.perimeter: ", g.perimeter())
}

func main() {
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
