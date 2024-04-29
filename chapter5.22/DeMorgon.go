package main

import "fmt"

func main() {
	x := 3
	y := 8
	a := 10
	b := 10
	g := 5
	i := 2
	j := 7

	originalA := !(x < 5) && !(y >= 7)
	originalB := !(a == b) || !(g != 5)
	originalC := !((x <= 8) && (y > 4))
	originalD := !((i > 4) || (j <= 6))

	newA := (x >= 5) && (y < 7)
	newB := (a != b) || (g == 5)
	newC := (x > 8) || (y <= 4)
	newD := (i <= 4) && (j > 6)

	fmt.Printf("Original expression a: %v\n", originalA)
	fmt.Printf("New expression a: %v\n", newA)

	fmt.Printf("Original expression b: %v\n", originalB)
	fmt.Printf("New expression b: %v\n", newB)

	fmt.Printf("Original expression c: %v\n", originalC)
	fmt.Printf("New expression c: %v\n", newC)

	fmt.Printf("Original expression d: %v\n", originalD)
	fmt.Printf("New expression d: %v\n", newD)
}
