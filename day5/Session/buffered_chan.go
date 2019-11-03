package main

import (
	"fmt"
	"time"
)

func buffChan(c <-chan int) {
	for i := range c {
		printVal(i)
	}
}

func printVal(i int) {
	fmt.Println("Task : ", i)
}

func main() {
	fmt.Println("Inside main")
	c := make(chan int, 6)
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println(cap(c), len(c))
	go buffChan(c)
	time.Sleep(time.Second)
}
