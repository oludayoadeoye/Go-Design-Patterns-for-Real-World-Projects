package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstNumber string
	var secondNumber string

	var firstNumInt int
	var secondNumInt int

	fmt.Print("Enter the first gifted number: ") // user prompt
	fmt.Scan(&firstNumber)                       // store input

	fmt.Print("Enter the second gifted number: ")
	fmt.Scan(&secondNumber)

	firstNumInt, _ = strconv.Atoi(firstNumber) // convert to int
	secondNumInt, _ = strconv.Atoi(secondNumber)

	fmt.Println(firstNumInt + secondNumInt) // addition of two strings
}
