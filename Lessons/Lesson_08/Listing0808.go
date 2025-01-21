package main

import (
	"fmt"
)

// this function accepts a variable number of input values
func sumN(numbers ...int) {
	sum := 0
	for i, num := range numbers {
		// display values to the user
		fmt.Println("Current element:", num, "; Current index:", i)
		sum += num
	}

	// print sum of all input values
	fmt.Println("Sum of values:", sum)
	return
}

func multiplyN(numbers ...int) {
	product := 1
	for n, num := range numbers {
		fmt.Println("Current element:", num, "; Current index:", n)
		product *= num
	}

	// print product of all input values
	fmt.Println("Product of values:", product)
	return
}

func main() {
	sumN(4, 6, 5)
	sumN(4, 6, 5, 6, 7, 8)

	multiplyN(4, 6, 5)
	multiplyN(4, 6, 5, 6, 7, 8)
}
