package main

import "fmt"

// main is the entry point of the program.
// It demonstrates the usage of arrays in Go, including declaration, initialization,
// accessing elements, and iterating over multi-dimensional arrays.
func main() {

	// Declare an array of 5 integers and print its initial state.
	var a [5]int
	fmt.Println("emp: ", a)

	// Set the value of the 5th element and print the array.
	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])
	fmt.Println("len: ", len(a))

	// Declare and initialize an array with specific values.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	// Declare and initialize an array using the ellipsis syntax.
	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", c)

	// Initialize an array with specific indices and values.
	// Elements in between will be zeroed.
	c = [...]int{3: 200, 500}
	fmt.Println("idx", c)
	c = [...]int{100, 3: 200, 500}
	fmt.Println("idx", c)

	// Declare a 2x3 multi-dimensional array and populate it using nested loops.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2b: ", twoD)

	// Initialize a 2x3 multi-dimensional array with specific values.
	twoD = [2][3]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println("2b: ", twoD)
}
