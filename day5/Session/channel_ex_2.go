package main

import "fmt"

func main() {
	c := make(chan int)
	d := make(chan int)
	fmt.Println("Inside main")
	go withTwoChannels(c, d, 1)
	fmt.Println(<-c)
	fmt.Println(<-d)
}

func withTwoChannels(c, d chan int, i int) {
	fmt.Println("Inside two channels")
	c <- i
	d <- i
}
