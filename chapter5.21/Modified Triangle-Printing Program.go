package main

import "fmt"

func main() {
	fmt.Println("Combined Triangle Patterns:")

	for row := 1; row <= 10; row++ {

		for col := 1; col <= row; col++ {
			fmt.Print("*")
		}

		for col := 1; col <= 10-row; col++ {
			fmt.Print(" ")
		}

		for col := 1; col <= 10-row; col++ {
			fmt.Print(" ")
		}
		for col := 1; col <= row; col++ {
			fmt.Print("*")
		}

		for col := 1; col <= 2; col++ {
			fmt.Print(" ")
		}

		for col := 1; col <= 10-row; col++ {
			fmt.Print(" ")
		}
		for col := 1; col <= row; col++ {
			fmt.Print("*")
		}
		for col := row - 1; col >= 1; col-- {
			fmt.Print("*")
		}

		for col := 1; col <= 2; col++ {
			fmt.Print(" ")
		}

		for col := 1; col <= row-1; col++ {
			fmt.Print(" ")
		}
		for col := 1; col <= 10-row+1; col++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
