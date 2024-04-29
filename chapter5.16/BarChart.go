package main

import (
	"fmt"
)

func main() {
	var numbers [5]int

	for i := 0; i < 5; i++ {
		fmt.Printf("Enter number %d (between 1 and 30): ", i+1)
		fmt.Scanln(&numbers[i])

		if numbers[i] < 1 || numbers[i] > 30 {
			fmt.Println("Invalid input. Please enter a number between 1 and 30.")
			i--
		}
	}

	for _, num := range numbers {
		for i := 0; i < num; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
