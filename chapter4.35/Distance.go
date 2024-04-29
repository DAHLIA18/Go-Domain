package main

import "fmt"

func main() {
	var x1, y1, x2, y2 int

	fmt.Println("Enter the coordinates of the first point:")
	fmt.Print("x1: ")
	fmt.Scanln(&x1)
	fmt.Print("y1: ")
	fmt.Scanln(&y1)

	fmt.Println("Enter the coordinates of the second point:")
	fmt.Print("x2: ")
	fmt.Scanln(&x2)
	fmt.Print("y2: ")
	fmt.Scanln(&y2)

	if x1 == x2 {
		fmt.Println("The points are located on a line perpendicular to the y-axis.")
	} else if y1 == y2 {
		fmt.Println("The points are located on a line perpendicular to the x-axis.")
	} else {
		fmt.Println("The points are not located on a line perpendicular to an axis.")
	}
}
