package main

import (
	"fmt"
)

func main() {
	var baseSalary float64 = 200.0
	var commissionRate float64 = 0.09
	var totalSales float64 = 0.0

	fmt.Print("Enter the number of items sold: ")
	var numItems int
	fmt.Scanln(&numItems)

	for i := 0; i < numItems; i++ {
		fmt.Printf("Enter the value of item %d: ", i+1)
		var itemValue float64
		fmt.Scanln(&itemValue)
		totalSales += itemValue
	}

	commission := totalSales * commissionRate
	totalEarnings := baseSalary + commission

	fmt.Printf("Total sales: $%.2f\n", totalSales)
	fmt.Printf("Commission: $%.2f\n", commission)
	fmt.Printf("Total earnings: $%.2f\n", totalEarnings)
}
