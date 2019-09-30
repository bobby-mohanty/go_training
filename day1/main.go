package main

import "fmt"

func main() {
	// Fibonacci Series
	printFib()
	fmt.Println()

	// Reverse a number
	reverseNum()
	fmt.Println()

	// Print Triangle
	printTriangle()
	fmt.Println()

	// Print Reverse Triangle
	printRevTriangle()
	fmt.Println()

	// Print number traingle
	printNumTriangle()
	fmt.Println()

	// Print rhombus
	printRhombus()
	fmt.Println()
}

func printFib() {
	var n int
	fmt.Println("Fibonacci Series \nEnter number of terms to print: ")
	fmt.Scanf("%d\n", &n)
	f := 0
	s := 1
	t := f + s
	fmt.Print(f, ", ", s)
	for i := 2; i < n; i++ {
		t = f + s
		fmt.Print(", ", t)
		f = s
		s = t
	}
	fmt.Print("\n")
}

func reverseNum() {
	var num int
	fmt.Println("Reverse a number \nEnter number:")
	fmt.Scanf("%d\n", &num)
	res := 0
	number := num
	for number > 0 {
		res = (res * 10) + (number % 10)
		number = number / 10
	}
	fmt.Printf("Original No.: %d, Reveresed No.: %d\n", num, res)
}

func printTriangle() {
	var n int
	fmt.Println("Print Forward Triangle \nEnter max no.: ")
	fmt.Scanf("%d\n", &n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func printRevTriangle() {
	var n int
	fmt.Println("Print Reverse Triangle \nEnter max no.: ")
	fmt.Scanf("%d\n", &n)
	for i := 1; i <= n; i++ {
		for k := 1; k <= n-i; k++ {
			fmt.Print("  ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func printNumTriangle() {
	var n int
	num := 0
	fmt.Println("Print Number Triangle \nEnter max levels: ")
	fmt.Scanf("%d\n", &n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			num++
			fmt.Printf("%d ", num)
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func printRhombus() {
	var num int
	fmt.Println("Print rhombus. \nEnter levels: ")
	fmt.Scanf("%d\n", &num)
	for i := 1; i <= num; i++ {
		for j := 1; j <= num-i; j++ {
			fmt.Print("  ")
		}
		for k := 1; k < (2 * i); k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("  ")
		}
		for k := 1; k < 2*(num-i); k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	fmt.Print("\n")
}
