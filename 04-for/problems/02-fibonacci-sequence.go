package main

import "fmt"

// main is the entry point of the program.
// It generates and prints the Fibonacci sequence up to a specified limit.
func main() {
	// Define the upper limit for the Fibonacci sequence.
	end := 1000

	// Initialize the Fibonacci sequence with the first two numbers.
	fib := []int{0, 1}

	// Generate the Fibonacci sequence until the next number exceeds the limit.
	for {
		// Calculate the next number in the sequence.
		fibLen := len(fib)
		nextNum := fib[fibLen-1] + fib[fibLen-2]

		// Break the loop if the next number exceeds the limit.
		if nextNum > end {
			break
		}

		// Append the next number to the sequence.
		fib = append(fib, nextNum)
	}

	// Print the generated Fibonacci sequence.
	fmt.Printf("Fibonacci Sequence until %d is %v\n", end, fib)
}
