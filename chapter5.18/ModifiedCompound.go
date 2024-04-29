package main

import "fmt"

func main() {
	const principal = 100000
	rate := 5
	var amount int

	fmt.Printf("%4s%21s\n", "Year", "Amount on deposit")

	for year := 1; year <= 10; year++ {

		amount = principal + (principal*rate)/100

		dollars := amount / 100
		cents := amount % 100
		fmt.Printf("%4d%19d.%02d\n", year, dollars, cents)

		principal = amount
	}
}
