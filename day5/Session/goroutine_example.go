package main

import (
	"fmt"
	"runtime"
	"time"
)

func some(n int) {
	fmt.Println("B -> ", n)
}

func main() {
	fmt.Println("A")
	for i := 0; i < 5; i++ {
		go some(i)
	}
	time.Sleep(time.Second * 3)
	fmt.Println("C")
	fmt.Println(runtime.NumGoroutine())
}
