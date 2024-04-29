package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	var n int
	fmt.Print("Enter a nonNegative integer: ")
	fmt.Scanln(&n)
	if n < 0 {
		fmt.Println("Invalid input. Please enter a nonnegative integer.")
		return
	}
	fmt.Printf("%d! = %d\n", n, factorial(n))
}
