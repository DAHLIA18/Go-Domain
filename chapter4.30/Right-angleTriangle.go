package main

import "fmt"

func main() {
	var baseLength int

	for {
		fmt.Print("Enter the length of the base of the triangle (1-10): ")
		_, err := fmt.Scanln(&baseLength)
		if err != nil || baseLength < 1 || baseLength > 10 {
			fmt.Println("Invalid input. Please enter a number between 1 and 10.")
			continue
		}
		break
	}

	for i := 1; i <= baseLength; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
