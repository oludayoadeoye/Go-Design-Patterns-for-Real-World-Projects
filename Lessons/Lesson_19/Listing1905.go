package main

import (
   "fmt"
   "sort"
)

func main() {
   // define a slice
   numbers := []int{67, 18, 62, 60, 25, 64, 75, 5, 17, 55}
   fmt.Println("Original Numbers:", numbers)

   sort.Ints(numbers)
   fmt.Println("Sorted Numbers:", numbers)

   sort.Sort(sort.Reverse(sort.IntSlice(numbers))) 
   fmt.Println("Sorted Numbers:", numbers)
}
