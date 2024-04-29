package main

import "fmt"

func main() {
	var targetSum int

	fmt.Print("Enter the target sum: ")

	_, err := fmt.Scanln(&targetSum)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	sum := 0

	for sum < targetSum {
		var num int

		fmt.Print("Enter an integer: ")
		_, err := fmt.Scanln(&num)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		sum += num
	}

	fmt.Printf("Target sum %d has been reached or exceeded.\n", targetSum)
}
