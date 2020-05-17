package main

import (
	"fmt"
	"math"
)

/*
same request format/endpoint/method
one interface
two different shapes
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
