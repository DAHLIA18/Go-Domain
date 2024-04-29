package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 5; j++ {
			if (i%2 == 1 && j%2 == 1) || (i%2 == 0 && j%2 == 0) {
				fmt.Print("@")
			} else if (i%2 == 1 && j%2 == 0) || (i%2 == 0 && j%2 == 1) {
				fmt.Print("$")
			}
		}
		fmt.Println()
	}
}
