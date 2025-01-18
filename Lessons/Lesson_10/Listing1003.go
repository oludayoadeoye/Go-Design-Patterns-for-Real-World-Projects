package main

import "fmt"

func main() {
   var a int = 20   // a stores the value 20
   var b *int = &a  // b stores the memory address of variable a

   fmt.Println(a)
   fmt.Println(b)
}