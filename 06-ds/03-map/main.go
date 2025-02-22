package main

import (
	"fmt"
	"maps"
	"reflect"
	"unsafe"
)

// main is the entry point of the program.
// It demonstrates the usage of maps in Go, including creation, insertion, deletion,
// clearing, and comparison of maps.
func main() {
	// Create a map m with string keys and int values.
	m := make(map[string]int)

	// Insert key-value pairs into the map m.
	m["Egypt"] = 110
	m["USA"] = 360
	m["UK"] = 90
	printVar(m)

	// Delete the key "USA" from the map m.
	delete(m, "USA")
	printVar(m)

	// Clear all elements from the map m.
	clear(m)
	printVar(m)

	// Create and initialize a map n with key-value pairs.
	n := map[string]int{"Egypt": 120, "USA": 400}
	printVar(n)

	// Create and initialize another map n2 and compare it with n.
	n2 := map[string]int{"Egypt": 120, "USA": 400}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}

// printVar prints the details of the given element, including its length, size, and type.
// element can be of any type.
func printVar(element interface{}) {
	// Determine the type of the element
	value := reflect.ValueOf(element)
	typ := reflect.TypeOf(element)

	// Print the details of the element
	fmt.Printf("| Map elements: %v \n| len: %d\n| size: %d\n| type: %v\n=====================================================\n", element, value.Len(), unsafe.Sizeof(element), typ)
}
