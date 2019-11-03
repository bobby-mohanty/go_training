package main

import "fmt"

func something(c chan bool) {
	fmt.Println("Inside go routine")
	c <- true
}

func main() {
	fmt.Println("Inside main")
	c := make(chan bool)
	go something(c)
	fmt.Println("->", <-c)
}
