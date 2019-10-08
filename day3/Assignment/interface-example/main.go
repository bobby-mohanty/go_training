package main

import (
	"fmt"
	"math"
)

type shape interface {
	getName() string
	area() float64
	perimeter() float64
}

type square struct {
	side float64
	name string
}

type rectangle struct {
	l float64
	b float64
	name string
}

type circle struct {
	radius float64
	name string
}

func display(s shape) {
	fmt.Printf("Area of the shape %s is %0.2f\n", s.getName(), s.area())
	fmt.Printf("Perimeter of the shape %s is %0.2f\n", s.getName(), s.perimeter())
}

func (r rectangle) area() float64 {
	return float64(r.l * r.b)
}
func (s square) area() float64 {
	return float64(s.side * s.side)
}
func (c circle) area() float64 {
	return float64(math.Pi * c.radius * c.radius)
}

func (r rectangle) perimeter() float64 {
	return float64(r.l*2 + 2*r.b)
}
func (s square) perimeter() float64 {
	return float64(4 * s.side)
}
func (c circle) perimeter() float64 {
	return float64(2 * math.Pi * c.radius)
}

func (r rectangle) getName() string {
	return r.name
}
func (s square) getName() string {
	return s.name
}
func (c circle) getName() string {
	return c.name
}

func main() {
	c := circle{
		radius: 5,
		name:   "CIRCLE",
	}
	r := rectangle{
		l:    5,
		b:    3,
		name: "RECTANGLE",
	}
	s := square{
		side: 6,
		name: "SQUARE",
	}
	display(c)
	display(r)
	display(s)
}
