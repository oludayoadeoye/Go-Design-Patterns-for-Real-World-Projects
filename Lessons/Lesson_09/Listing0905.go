package main

import "fmt"

func main() {
   // create an array of three integers
   var numbers [3]int

   // assign a value to each position in the array
   numbers[0] = 1
   numbers[1] = 34
   numbers[2] = 3455
   numbers[0] = 999
   numbers[0] = 50000
   numbers[1] = numbers[2]

   // display the array
   fmt.Println(numbers)
}