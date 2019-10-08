package main

import (
	"fmt"
	"os"
)

func copyIt() error {
	src, err := os.Open("day45.txt")
	defer func() {
		recover()
		fmt.Println("recovered from panic")
	}()
	if err != nil {
		panic("new error")
	}
	b := make([]byte, 50)
	count, _ := src.Read(b)
	fmt.Printf("Bytes read : %d \n\n Text File : %s", count, string(b))
	return nil
}

func main() {
	copyIt()
}
