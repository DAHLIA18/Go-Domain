package main

import "fmt"

func main() {
	total := 0
	gradeCounter := 1

	for gradeCounter <= 10 {
		fmt.Print("Enter grade (1 or 2): ")
		var grade int
		_, err := fmt.Scanln(&grade)
		if err != nil || (grade != 1 && grade != 2) {
			fmt.Println("Invalid input. Please enter 1 or 2.")
			continue
		}

		total += grade
		gradeCounter++
	}

	average := float64(total) / 10
	fmt.Printf("Class average is %.2f\n", average)
}
