package main

import "fmt"

// main is the entry point of the program.
// It demonstrates basic operations in Go including string concatenation,
// arithmetic operations, and boolean logic.
func main() {
	// Concatenate and print two strings.
	fmt.Println("go" + "lang")

	// Perform and print the result of an integer addition.
	fmt.Println("1 + 1 =", 1+1)

	// Perform and print the result of a floating-point division.
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Perform and print the result of an integer division.
	fmt.Println("7/3 =", 7/3)

	// Perform and print the result of a boolean AND operation.
	fmt.Println("true && false =", true && false)

	// Perform and print the result of a boolean OR operation.
	fmt.Println("true || false =", true || false)

	// Perform and print the result of a boolean NOT operation.
	fmt.Println("!true =", !true)
}
