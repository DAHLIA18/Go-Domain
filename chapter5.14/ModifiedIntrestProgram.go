package main

import (
	"fmt"
)

func main() {
	var principal float64 = 1000.0
	rate := 0.05
	years := 10

	fmt.Printf("%4s%21s\n", "Year", "Amount on deposit")

	for rate <= 0.10 {
		fmt.Printf("Interest rate: %.2f%%\n", rate*100)

		for year := 1; year <= years; year++ {
			amount := principal * pow(1.0+rate, float64(year))
			fmt.Printf("%4d%21.2f\n", year, amount)
		}

		rate += 0.01
		fmt.Println()
	}
}

func pow(base, exponent float64) float64 {
	result := 1.0

	for i := 0; i < int(exponent); i++ {
		result *= base
	}

	return result
}
