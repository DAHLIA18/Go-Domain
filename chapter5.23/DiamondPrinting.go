package main

import "fmt"

func main() {
	size := 7

	for i := 1; i <= size; i += 2 {
		printSpaces((size - i) / 2)
		printAsterisks(i)
		fmt.Println()
	}

	for i := size - 2; i >= 1; i -= 2 {
		printSpaces((size - i) / 2)
		printAsterisks(i)
		fmt.Println()
	}
}

// Function to print spaces
func printSpaces(count int) {
	for i := 0; i < count; i++ {
		fmt.Print(" ")
	}
}

func printAsterisks(count int) {
	for i := 0; i < count; i++ {
		fmt.Print("*")
	}
}
