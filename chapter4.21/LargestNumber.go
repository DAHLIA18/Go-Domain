package main

import (
	"fmt"
	"math"
)

func main() {
	var counter int
	var largest int = math.MinInt64

	for counter < 10 {
		var number int
		fmt.Print("Enter an integer: ")
		fmt.Scanln(&number)

		if number > largest {
			largest = number
		}

		counter++
	}

	fmt.Println("The largest number is:", largest)
}
