package main

import "fmt"

// main is the entry point of the program.
// It demonstrates two methods to sum all even numbers between 1 and 100.
func main() {
	//===============================================================
	// Write a program that uses a for loop to sum all even numbers between 1 and 100.
	//===============================================================
	// 01- using For Loop
	//===============================================================

	// Initialize the sum variable to store the sum of even numbers.
	var sum int = 0
	// Define the start and end range.
	var start int = 1
	var end int = 100

	// Iterate from start to end.
	for i := start; i <= end; i++ {
		// Check if the number is even.
		if i%2 == 0 {
			// Add the even number to the sum.
			sum += i
		}
	}
	// Print the sum of all even numbers between start and end.
	fmt.Printf("Sum all even numbers between %d and %d is %d\n", start, end, sum)

	//===============================================================
	// 02- using Mathematics equations
	//===============================================================
	// Redefine the start and end range for the mathematical approach.
	start = 2
	end = 100
	// Calculate the number of even pairs.
	numOfPairs := (100-2)/2 + 1
	// Calculate the sum using the arithmetic series formula.
	sum = (100 + 2) * (numOfPairs / 2)
	// Print the sum of all even numbers between start and end.
	fmt.Printf("Sum all even numbers between %d and %d is %d", start, end, sum)
}
