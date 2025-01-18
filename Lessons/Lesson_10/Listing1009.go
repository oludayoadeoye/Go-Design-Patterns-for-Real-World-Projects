ppackage main

import "fmt"

func main() {
   var b *int   // create a pointer variable b
   fmt.Println(b) 
   if b == nil {
      fmt.Println("b is nil")
   } else {
      fmt.Println(*b)
   }
}