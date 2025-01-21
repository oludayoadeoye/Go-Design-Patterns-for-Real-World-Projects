package main

import (
   "fmt"
)

func main() {
   // create an empty slice
   var s []int
   fmt.Println(s)
   s = append(s, 10)
   fmt.Println(s)
   s = append(s, 11)
   s = append(s, 12)
   s = append(s, 13)
   fmt.Println(s)
}