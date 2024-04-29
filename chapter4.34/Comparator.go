package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Print("Enter the first number: ")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print("Enter the second number: ")
	_, err = fmt.Scanln(&num2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if num1 == num2 {
		fmt.Println("0")
	} else if num1 > num2 {
		fmt.Println("1")
	} else {
		fmt.Println("-1")
	}
}
