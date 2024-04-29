package main

import (
	"fmt"
)

func main() {
	var number int

	for {
		fmt.Print("Enter a five-digit integer: ")
		_, err := fmt.Scanln(&number)
		if err != nil || number < 10000 || number > 99999 {
			fmt.Println("Invalid input. Please enter a five-digit integer.")
			continue
		}
		break
	}

	originalNumber := number
	reversedNumber := 0

	for number != 0 {
		digit := number % 10
		reversedNumber = reversedNumber*10 + digit
		number /= 10
	}

	if originalNumber == reversedNumber {
		fmt.Println("The number is a palindrome.")
	} else {
		fmt.Println("The number is not a palindrome.")
	}
}
