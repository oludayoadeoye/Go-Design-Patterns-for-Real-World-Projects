package main

import "fmt"

func main() {
   // create a simple array
   numbers := []int{100,1000,10000} 
   // print each value in the array
   var ctr int
   for ctr = 0; ctr < len(numbers); ctr++ {
      fmt.Println(numbers[ctr])
   }

   // create an array of pointers
   var numbersptr [3] *int;

   // assign a pointer to each value in the original array and 
   // store them in the new array
   for ctr := 0; ctr < len(numbersptr); ctr ++ {
      numbersptr[ctr] = &numbers[ctr]
   }

   fmt.Println(numbersptr) // print the pointer array

   // print the values the pointers point to
   fmt.Println(*numbersptr[0]) 
   fmt.Println(*numbersptr[1])
   fmt.Println(*numbersptr[2])
}