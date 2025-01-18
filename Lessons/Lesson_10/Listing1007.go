ppackage main

import "fmt"

func main() {
   var myVar int = 20
   b := &myVar

   fmt.Println(b)    // print the memory address
   fmt.Println(*b)   // print the value stored in the memory address
}
