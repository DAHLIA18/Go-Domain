package main

import "fmt"

func main() {
	var taxRateLow = 0.15
	var taxRateHigh = 0.20

	var numCitizens int
	fmt.Print("Enter the number of citizens: ")
	fmt.Scanln(&numCitizens)

	for i := 0; i < numCitizens; i++ {
		fmt.Printf("\nCitizen %d\n", i+1)
		fmt.Print("Enter name: ")
		var name string
		fmt.Scanln(&name)
		fmt.Print("Enter earnings: ")
		var earnings float64
		fmt.Scanln(&earnings)

		var tax float64
		if earnings <= 30000 {
			tax = earnings * taxRateLow
		} else {
			tax = (30000 * taxRateLow) + ((earnings - 30000) * taxRateHigh)
		}

		fmt.Printf("Total tax for %s: $%.2f\n", name, tax)
	}
}
