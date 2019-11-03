package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for {
			fmt.Println(<-c)
			fmt.Println("Task")
		}
	}()
	c <- 5
	close(c)
}
