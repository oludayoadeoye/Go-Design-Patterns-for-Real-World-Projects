package main

import "fmt"

func main() {
   var a int = 20
   b := &a
   c := &a
   var x *int

   fmt.Print("Is c == b? ")
   fmt.Println(c == b) // compare c and b
 
   fmt.Print("Is x == b? ")
   fmt.Println(x == b) // compare x and b
}