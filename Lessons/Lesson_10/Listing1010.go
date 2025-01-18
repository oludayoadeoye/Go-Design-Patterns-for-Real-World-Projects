package main

import "fmt"

func main() {
   var a int = 20
   b := &a

   fmt.Print("The value stored in a: ")
   fmt.Println(a)  // print a

   fmt.Print("Memory address: ")
   fmt.Println(b)  // print the memory address

   fmt.Print("Value stored in a (via pointer): ")
   fmt.Println(*b) // print the value stored at the memory address

   *b = 30   // use the *b to change the value stored at the
             // memory address (or in variable a)
   fmt.Print("New value of a: ")
   fmt.Println(a)  // the value of a changed
}