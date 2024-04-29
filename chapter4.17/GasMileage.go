package main

import (
	"fmt"
)

func main() {
	var totalMiles, totalGallons int

	fmt.Println("Gas Mileage Calculator")
	fmt.Println("----------------------")

	for {
		var milesDriven, gallonsUsed int

		fmt.Print("Enter miles driven (or -1 to exit): ")
		_, err := fmt.Scanln(&milesDriven)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if milesDriven == -1 {
			break
		}

		fmt.Print("Enter gallons used: ")
		_, err = fmt.Scanln(&gallonsUsed)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		mpg := float64(milesDriven) / float64(gallonsUsed)
		totalMiles += milesDriven
		totalGallons += gallonsUsed

		fmt.Printf("Miles per gallon for this trip: %.2f\n", mpg)
		if totalGallons != 0 {
			combinedMPG := float64(totalMiles) / float64(totalGallons)
			fmt.Printf("Combined miles per gallon: %.2f\n", combinedMPG)
		}
	}

	fmt.Println("Exiting Gas Mileage Calculator.")
}
