package main

import (
	"fmt"
)

func main() {
	var binary int

	for {
		fmt.Print("Enter a binary integer (containing only 0s and 1s): ")
		_, err := fmt.Scanln(&binary)
		if err != nil || !isValidBinary(binary) {
			fmt.Println("Invalid input. Please enter a binary integer containing only 0s and 1s.")
			continue
		}
		break
	}

	decimal := 0
	position := 1

	for binary > 0 {
		digit := binary % 10
		decimal += digit * position
		binary /= 10
		position *= 2
	}

	fmt.Println("Decimal equivalent:", decimal)
}

func isValidBinary(n int) bool {
	for n > 0 {
		digit := n % 10
		if digit != 0 && digit != 1 {
			return false
		}
		n /= 10
	}
	return true
}
