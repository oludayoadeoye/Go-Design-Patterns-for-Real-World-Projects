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
   fruits := []string{"pear", "pineapple", "mango", "banana", "fig"}
   sort.Sort(sort.Reverse(sort.IntSlice(numbers))) 
   sort.Sort(sort.Reverse(sort.StringSlice(fruits))) 
   fmt.Println("Sorted Numbers:", numbers)
   fmt.Println("REVERSED FRUITS:", fruits)
}
