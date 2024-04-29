package main

import (
	"fmt"
)

func main() {
	var count int
	fmt.Print("Enter the number of values: ")
	fmt.Scanln(&count)

	if count <= 0 {
		fmt.Println("Invalid input. Please enter a positive integer.")
		return
	}

	var minValue, maxValue, value, sum int

	for i := 1; i <= count; i++ {
		fmt.Printf("Enter value %d: ", i)
		fmt.Scanln(&value)

		if i == 1 {
			minValue = value
			maxValue = value
		} else {
			if value < minValue {
				minValue = value
			}
			if value > maxValue {
				maxValue = value
			}
		}
	}

	sum = minValue + maxValue
	fmt.Printf("Minimum value: %d\n", minValue)
	fmt.Printf("Maximum value: %d\n", maxValue)
	fmt.Printf("Sum of extremes: %d\n", sum)
}
