package main

import "fmt"

func main() {
	fmt.Println("Pythagorean Triples (side1, side2, hypotenuse):")

	for side1 := 1; side1 <= 500; side1++ {

		for side2 := 1; side2 <= 500; side2++ {

			for hypotenuse := 1; hypotenuse <= 500; hypotenuse++ {

				if isPythagoreanTriple(side1, side2, hypotenuse) {
					fmt.Printf("(%d, %d, %d)\n", side1, side2, hypotenuse)
				}
			}
		}
	}
}
func isPythagoreanTriple(side1, side2, hypotenuse int) bool {
	return side1*side1+side2*side2 == hypotenuse*hypotenuse
}
