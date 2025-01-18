package main

import "fmt"

func main() {
   // create an array of three integers
   var number int 
   var numbers [3]int

   // assign a value to the basic variable 
   number = "one"

   // assign a value to each position in the array
   numbers[0] = "one"
   numbers[1] = "two" 
   numbers[2] = "three"

   // display the int
   fmt.Println(number)
   // display the array
   fmt.Println(numbers)
}