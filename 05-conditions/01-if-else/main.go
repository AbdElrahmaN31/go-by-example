package main

import (
	"fmt"
)

// main is the entry point of the program.
// It demonstrates the usage of if-else statements in Go to check conditions
// and print appropriate messages based on those conditions.
func main() {

	// Check if 7 is even or odd.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Check if age is less than 30.
	if age := 10; age < 30 {
		fmt.Println("Age is less than or equal to 30")
	}

	// Check if num is negative, has one digit, or has multiple digits.
	if num := 9; num < 0 {
		fmt.Println("Num is negative")
	} else if num < 10 {
		fmt.Println("Num has one digit")
	} else {
		fmt.Println("Num has multiple digits")
	}

	// Check if num is greater than 0 and less than 10.
	if num := 9; num > 0 && num < 10 {
		fmt.Println("Num is greater than 0 and less than 10")
	}
}
