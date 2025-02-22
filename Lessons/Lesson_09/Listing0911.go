package main

import "fmt"

func main() {
	// declare a two-dimensional array with two sizes instead of one
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)
	fmt.Println(matrix[0][0])
	fmt.Println(matrix[0][1])
	fmt.Println(matrix[0][2])
	fmt.Println(matrix[1][0])
	fmt.Println(matrix[1][1])
	fmt.Println(matrix[1][2])
	fmt.Println(matrix[2][1])
	fmt.Println(matrix[2][0])
	fmt.Println(matrix[2][2])
}
