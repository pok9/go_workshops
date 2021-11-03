package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	r := rectangle{width: 10, height: 10}
	c := circle{radius: 10}

	fmt.Printf("Rectangle: %f\n", getArea(r))
	fmt.Printf("Circle: %f\n", getArea(c))

	showInfo(r)
	showInfo(c)

	castToRectangle(c) //Error
	castToRectangle(r) //No Error
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func getArea(s shape) float64 {
	return s.area()
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func showInfo(s shape) {
	t := reflect.TypeOf(s).Name()
	switch t {
	case "rectangle":
		r := s.(rectangle)
		fmt.Printf("Rect width: %f, height: %f\n", r.width, r.height)
	case "circle":
		c := s.(circle)
		fmt.Printf("Circle radius: %f\n", c.radius)
	}
}

func castToRectangle(s shape) {
	r, ok := s.(rectangle)
	if !ok {
		fmt.Println("Casting Error")
	} else {
		fmt.Printf("Casting to Rectangle Success w: %f,h: %f\n", r.width, r.height)
		fmt.Printf("Casting to Rectangle Success Area %f", getArea(s))
	}
}
