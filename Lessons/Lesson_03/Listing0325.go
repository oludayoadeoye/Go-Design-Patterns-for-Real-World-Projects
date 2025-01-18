package main

import "fmt"

func main() {
	// firstname := "John"
	// street_address := "123 Main St"
	// birth_year := 2000

	var former_street_adress string
	var new_strt_address string

	fmt.Println("Enter your former street address: ")
	fmt.Scanln(&former_street_adress)

	fmt.Println("Enter your new street address: ")
	fmt.Scanln(&new_strt_address)

	fmt.Println("now , I get it, you used to live in ", former_street_adress, " but you now live here, ", new_strt_address, " right?")
}
