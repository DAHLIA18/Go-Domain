package main

import (
	"fmt"
	"math"
)

func main() {
	terms := 200000
	piApprox := 0.0
	sign := 1.0

	for i := 1; i <= terms; i += 2 {
		piApprox += sign * (1.0 / float64(i))
		sign *= -1
	}

	piApprox *= 4
	fmt.Printf("Approximation of π using %d terms: %.10f\n", terms, piApprox)

	termsForAccuracy := findTermsForAccuracy(3.14159)
	fmt.Printf("Terms required for accuracy: %d\n", termsForAccuracy)
}

func findTermsForAccuracy(desiredValue float64) int {
	terms := 0
	piApprox := 0.0
	sign := 1.0

	for {
		terms += 2
		piApprox += sign * (1.0 / float64(terms))
		sign *= -1

		piApprox *= 4
		if math.Floor(piApprox*1e5) == math.Floor(desiredValue*1e5) {
			break
		}
	}

	return terms
}
