package main

import "fmt"

type person struct {
	name string
	age  int
}

type some interface {
	display()
}

func (p person) display() {
	fmt.Println("something")
}

func main() {
	obj := person{"bobby", 12}
	obj.display()
}
