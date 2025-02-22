package main

import (
	"fmt"
	"slices"
)

// main is the entry point of the program.
// It demonstrates the usage of slices in Go, including initialization, appending elements,
// copying slices, slicing slices, and comparing slices.
func main() {
	// Declare an uninitialized slice of strings and print its state.
	var s []string
	fmt.Println("un-initialized: ", s, s == nil, len(s))

	// Initialize the slice with a length of 3 and print its state.
	s = make([]string, 3)
	fmt.Println("emp: ", s, s == nil, len(s))

	// Append elements to the slice and print the updated slice and its length.
	s = append(s, "a")
	s = append(s, "b")
	s = append(s, "d")
	s = append(s, "e")
	s = append(s, "f")
	s = append(s, "g")
	s = append(s, "h")
	fmt.Println("added: ", s, " | len: ", len(s))

	// Create a new slice c with the same length as s and copy elements from s to c.
	c := make([]string, len(s))
	fmt.Println("new c:", c, len(c))
	copy(c, s)
	fmt.Println("c after copy:", c, len(c))

	// Create a slice l from s[3:5] and print it.
	l := s[3:5]
	fmt.Println("slice from s[1:3]: ", l, " | len: ", len(l))

	// Create a slice l from s[3:] and print it.
	l = s[3:]
	fmt.Println("slice from s[3:]: ", l, " | len: ", len(l))

	// Declare and initialize a slice t with specific values.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Declare and initialize another slice t2 and compare it with t.
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Create a 2D slice of integers with a length of 3 and capacity of 10.
	twoD := make([][]int, 3, 10)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	// Print the 2D slice, its length, and its capacity.
	fmt.Printf("2d slice: %v\nIt's len is %d, and it's cap is %d", twoD, len(twoD), cap(twoD))

}
