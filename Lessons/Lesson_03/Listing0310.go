package main

import "fmt"

func main() {
	var message = "Hello, World!" // missing var
	email := "john@john.com"      // cannot use var and := together

	fmt.Println(message)
	fmt.Println(email)
}
