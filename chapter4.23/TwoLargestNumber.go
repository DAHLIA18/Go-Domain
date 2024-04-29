package main

import (
	"fmt"
	"math"
)

func main() {
	var largest, secondLargest int = math.MinInt64, math.MinInt64

	for i := 0; i < 10; i++ {
		var num int
		fmt.Printf("Enter number %d: ", i+1)
		fmt.Scanln(&num)

		if num > largest {
			secondLargest = largest
			largest = num
		} else if num > secondLargest {
			secondLargest = num
		}
	}

	fmt.Printf("Largest number: %d\n", largest)
	fmt.Printf("Second largest number: %d\n", secondLargest)
}
