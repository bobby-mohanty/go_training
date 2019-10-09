package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func readFile(filename string) {
	filename, _ = filepath.Abs(filename)
	fp, err := os.Open(filename)
	content := make([]byte, 50)
	defer func() {
		fmt.Println("Defer for closing file pointer fp")
		_ = fp.Close()
		recover()
	}()
	if err != nil {
		fmt.Printf("file '%s' doesn't exist\n", filename)
		panic(fmt.Sprintf("file '%s' doesn't exist\n", filename))
	} else {
		if _, e := fp.Read(content); e != nil {
			fmt.Println("Cannot read file content ")
			panic("Cannot read file content ")
		}
		fmt.Println("File Content :", string(content))
	}
}

func divide(num1, num2 int) (result float64, err error) {
	if num2 == 0 {
		return 0, errors.New("divide by zero")
	}
	return float64(num1 / num2), nil
}

func main() {
	defer fmt.Println("Defer in main function")
	readFile("test.txt")
	readFile("some.txt")
	result, err := divide(5, 2)
	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println("Result of 5/2:", result)
	}
	result, err = divide(5, 0)
	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println("Result of 5/0:", result)
	}
}
