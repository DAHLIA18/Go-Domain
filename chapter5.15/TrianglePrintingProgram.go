package main

import "fmt"

func main() {

	fmt.Println("(a)")
	for i := 1; i <= 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("\n(b)")
	for i := 10; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("\n(c)")
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("\n(d)")
	for i := 10; i >= 1; i-- {
		for j := 1; j <= 10-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
