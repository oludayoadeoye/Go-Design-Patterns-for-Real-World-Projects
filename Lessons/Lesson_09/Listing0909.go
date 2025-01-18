package main

import "fmt"

func main() {
   var numbers [20]int
   for x := 0; x < 20; x++ {
      numbers[x] = x
   }
   fmt.Println(numbers)
}
