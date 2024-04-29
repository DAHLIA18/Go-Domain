package main

import "fmt"

func main() {
	fmt.Println("The Twelve Days of Christmas:")

	days := [...]string{
		"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth",
	}

	gifts := [...]string{
		"a Partridge in a Pear Tree",
		"Two Turtle Doves",
		"Three French Hens",
		"Four Calling Birds",
		"Five Gold Rings",
		"Six Geese-a-Laying",
		"Seven Swans-a-Swimming",
		"Eight Maids-a-Milking",
		"Nine Ladies Dancing",
		"Ten Lords-a-Leaping",
		"Eleven Pipers Piping",
		"Twelve Drummers Drumming",
	}

	for i, day := range days {
		fmt.Printf("On the %s day of Christmas,\nMy true love gave to me:\n", day)
		for j := i; j >= 0; j-- {
			if j == 0 && i != 0 {
				fmt.Print("And ")
			}
			fmt.Println(gifts[j])
		}
		fmt.Println()
	}
}
