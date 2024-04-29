package main

import (
	"fmt"
)

func main() {
	fmt.Println("Item   Value")
	fmt.Println("1      239.99")
	fmt.Println("2      129.75")
	fmt.Println("3      99.95")
	fmt.Println("4      350.89")

	fmt.Println("\nN   N2   N3   N4")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%-4d%-5d%-5d%-5d\n", i, i*i, i*i*i, i*i*i*i)
	}
}
