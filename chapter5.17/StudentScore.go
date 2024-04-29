package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var studentGrades = map[string]string{
		"student 1": "",
		"student 2": "",
		"student 3": "",
		"student 4": "",
		"student 5": "",
	}

	reader := bufio.NewReader(os.Stdin)
	for student := range studentGrades {
		fmt.Printf("Enter grade for %s: ", student)
		grade, _ := reader.ReadString('\n')
		studentGrades[student] = strings.TrimSpace(grade)

		var gradeCounts = map[string]int{
			"A": 0,
			"B": 0,
			"C": 0,
			"D": 0,
		}

		for _, grade := range studentGrades {
			switch grade {
			case "A":
				gradeCounts["A"]++
			case "B":
				gradeCounts["B"]++
			case "C":
				gradeCounts["C"]++
			case "D":
				gradeCounts["D"]++
			default:
				fmt.Println("Invalid grade:", grade)
			}
		}

		fmt.Println("Grade counts:")
		for grade, count := range gradeCounts {
			fmt.Printf("%s: %d\n", grade, count)
		}
	}
}
