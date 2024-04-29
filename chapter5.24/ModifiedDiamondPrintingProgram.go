package main

import (
	"fmt"
)

func main() {
	var size int
	for {
		fmt.Print("Enter an odd number between 1 and 19: ")
		fmt.Scanln(&size)
		if size >= 1 && size <= 19 && size%2 == 1 {
			break
		}
		fmt.Println("Invalid input. Please enter an odd number between 1 and 19.")
	}

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
