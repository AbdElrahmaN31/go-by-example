package main

import (
	"fmt"
	"math"
)

// s is a constant string with a greeting message.
const s string = "Hello GOoOo!"

// main is the entry point of the program.
// It demonstrates the usage of constants, basic arithmetic operations,
// type conversion, and mathematical functions in Go.
func main() {
	// Print the constant string s.
	fmt.Println(s)

	// Declare a constant integer n.
	const n = 50000000000000000

	// Declare a constant d as the result of a division operation.
	const d = 3e20 / n
	// Print the value of d.
	fmt.Println(d)

	// Convert d to int16 and print the result.
	fmt.Println(int16(d))

	// Calculate and print the sine of n.
	fmt.Println(math.Sin(n))

	// Declare a constant Pi with the value of π.
	const Pi = 3.141592653589793

	// Declare a constant EarthRadius with the value of Earth's radius in kilometers.
	const EarthRadius = 6371

	// Print the value of Pi.
	println(Pi)

	// Print the value of EarthRadius.
	println(EarthRadius)
}
