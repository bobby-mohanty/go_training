package main

import (
	"fmt"
	"math"
)

type shape interface {
	getName() string
	area() float64
	perimeter() float64
	getDimensions() string
}

type circle struct {
	radius float64
	name string
}

func display(s shape) {
	fmt.Printf("Area of the shape %s is %0.2f\n", s.getName(), s.area())
	fmt.Printf("Perimeter of the shape %s is %0.2f\n", s.getName(), s.perimeter())
}

func (c circle) area() float64 {
	return float64(math.Pi * c.radius * c.radius)
}

func (c circle) perimeter() float64 {
	return float64(2 * math.Pi * c.radius)
}

func (c circle) getName() string {
	return c.name
}

func main() {
	c := circle{
		radius: 5,
		name:   "CIRCLE",
	}
	display(c)
}

// getDimension() is not implemented so could not use circle structure object
//
// ERROR :
//.\main.go:42:9: cannot use c (type circle) as type shape in argument to display:
//circle does not implement shape (missing getDimensions method)