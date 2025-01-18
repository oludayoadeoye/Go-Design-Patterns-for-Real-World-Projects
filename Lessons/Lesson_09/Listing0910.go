package main

import "fmt"

func main() {
 numbers := [4] int {1,3,5,7} 

   fmt.Println("Printing numbers:")
   fmt.Println(numbers)
   fmt.Println("Starting for loop...")
   for index, value := range numbers{
      fmt.Println(index)
      fmt.Println(value)
      fmt.Println("---")
   }
}