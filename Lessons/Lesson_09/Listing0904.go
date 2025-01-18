package main

import "fmt"

func main() {
   // create an array of three integers
   var numbers [3]int

   // assign a value to each position in the array
   numbers[0] = 1
   numbers[1] = 34
   numbers[2] = 3455
   numbers[3] = 30
   numbers[4] = 40
   numbers[99] = 990

   // display the array
   fmt.Println(numbers[0], numbers[1], numbers[2])
}