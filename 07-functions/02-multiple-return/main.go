package main

import "fmt"

// divide performs integer division and returns both the quotient and the remainder.
// Parameters:
// - dividend: The number to be divided.
// - divisor: The number by which to divide.
// Returns: The quotient and the remainder of the division.
func divide(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

func main() {
	// Example 1: Using both return values
	quotient, remainder := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder)

	// Example 2: Ignoring the remainder
	quotientOnly, _ := divide(20, 4)
	fmt.Printf("Quotient only: %d\n", quotientOnly)

	// Example 3: Ignoring the quotient
	_, remainderOnly := divide(15, 6)
	fmt.Printf("Remainder only: %d\n", remainderOnly)
}
