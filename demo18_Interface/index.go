package main

import (
	"fmt"
	"math"
)

func main() {
	r := rectangle{width: 10, height: 10}
	c := circle{radius: 10}

	fmt.Printf("Rectangle: %f\n", getArea(r))
	fmt.Printf("Circle: %f\n", getArea(c))
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func getArea(s shape) float64 {
	return s.area()
}

type shape interface {
	area() float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
