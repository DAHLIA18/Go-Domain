package main

import (
	"fmt"
)

func computeEx(x float64, numTerms int) float64 {
	ex := 1.0
	power := 1.0
	factorial := 1.0
	for i := 1; i <= numTerms; i++ {
		power *= x
		factorial *= float64(i)
		ex += power / factorial
	}
	return ex
}

func main() {
	var x float64
	var numTerms int
	fmt.Print("Enter the value of x: ")
	fmt.Scanln(&x)
	fmt.Print("Enter the number of terms to calculate e^x: ")
	fmt.Scanln(&numTerms)
	if numTerms < 1 {
		fmt.Println("Invalid input. Please enter a positive integer for the number of terms.")
		return
	}
	fmt.Printf("Value of e^%.2f with %d terms: %.4f\n", x, numTerms, computeEx(x, numTerms))
}
