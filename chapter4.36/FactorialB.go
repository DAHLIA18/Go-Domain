package main

import "fmt"

func estimateE(numTerms int) float64 {
	e := 1.0
	factorial := 1.0
	for i := 1; i <= numTerms; i++ {
		factorial *= float64(i)
		e += 1.0 / factorial
	}
	return e
}

func main() {
	var numTerms int
	fmt.Print("Enter the number of terms to calculate e: ")
	fmt.Scanln(&numTerms)
	if numTerms < 1 {
		fmt.Println("Invalid input. Please enter a positive integer.")
		return
	}
	fmt.Printf("Estimated value of e with %d terms: %f\n", numTerms, estimateE(numTerms))
}
