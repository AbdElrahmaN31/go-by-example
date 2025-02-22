package main

import "fmt"

// s is a constant string with a greeting message.
const s string = "Hello, world!"

// main is the entry point of the program.
// It demonstrates various types of for loops in Go, including traditional for loops,
// while-like loops, range-based loops, and infinite loops.
func main() {
	// Traditional For Loop
	for i := 0; i < 10; i++ {
		for j := 10; j > i; j-- {
			// Print the character '*' instead of its ASCII code.
			fmt.Printf("%c", '*')
		}
		println()
	}

	// Equal to while!
	j := 0
	for j < 10 {
		// Print the value of j.
		println(j)
		j++
	}

	// Range like in Python
	for i := range s {
		// Print each character in the string s.
		fmt.Printf("%c, ", s[i])
	}
	println()
	for i, c := range s {
		// Print the index and character in the string s.
		fmt.Printf("%d: %c, ", i, c)
	}
	println()

	// Range over a slice of integers from 0 to 5.
	for n := range 6 {
		// Print the even numbers.
		if n%2 == 0 {
			fmt.Println(n)
		}
	}

	// Infinite loop
	for {
		// Print a message and break the loop.
		println("infinite loop")
		break
	}
}
