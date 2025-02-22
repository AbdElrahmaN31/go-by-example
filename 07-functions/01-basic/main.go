package main

import "fmt"

// plus adds two integers and returns the result.
// Parameters:
// - num1: The first integer to add.
// - num2: The second integer to add.
// Returns: The sum of num1 and num2.
func plus(num1 int, num2 int) int {
	return num1 + num2
}

// plusPlus adds three integers and returns the result.
// Parameters:
// - num1: The first integer to add.
// - num2: The second integer to add.
// - num3: The third integer to add.
// Returns: The sum of num1, num2, and num3.
func plusPlus(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

// sum calculates the sum of any number of integers.
// Parameters:
// - nums: A variable number of integers to sum.
// Returns: The total sum of the provided integers.
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	// Call plus with 1 and 2, and print the result.
	result := plus(1, 2)
	fmt.Printf("1 + 2 = %d\n", result)

	// Call plusPlus with 1, 2, and 3, and print the result.
	result = plusPlus(1, 2, 3)
	fmt.Printf("1 + 2 + 3 = %d\n", result)

	// Call sum with multiple integers, and print the result.
	result = sum(1, 2, 3, 50, 60, 120)
	fmt.Printf("sum of (1, 2, 3, 50, 60, 120) is %d", result)
}
