package main

import "fmt"

func main() {
   numbers_1 := [3]int {5, 6, 7}
   fmt.Println(numbers_1)

   // copy the array to a new variable
   numbers_2 := numbers_1
   fmt.Println(numbers_2)

   // change a value in the new array
   numbers_2[0]=100
 
   // output both arrays
   fmt.Println(numbers_1)
   fmt.Println(numbers_2)
}