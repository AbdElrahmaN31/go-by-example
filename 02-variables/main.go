package main

import (
	"fmt"
)

// main is the entry point of the program.
// It demonstrates variable declaration, initialization, and printing in Go.
func main() {
	// Declare and initialize a string variable.
	var a = "Hello, world!"
	fmt.Println(a)

	// Declare and initialize two integer variables.
	var b, c int = 10, 5
	fmt.Println(b, c)

	// Declare and initialize a boolean variable.
	var d = true
	fmt.Println(d)

	// Declare and initialize a floating-point variable using short variable declaration.
	e := 13.14
	fmt.Println(e)
}
