package main

import "fmt"

func main() {
   var a int = 20  // a stores the value 20
   b := &a         // b stores the memory address of variable a

   var c float32 = 10.3 
   var d *float32 = &c

   var e string = "My string"
   var f *string = &e 

   var g uint = 42
   var h *uint = &g 

   fmt.Println(a)
   fmt.Println(b)
   fmt.Println("-------------")
   fmt.Println(c)
   fmt.Println(d)
   fmt.Println("-------------")
   fmt.Println(e)
   fmt.Println(f)
   fmt.Println("-------------")
   fmt.Println(g)
   fmt.Println(h)
   fmt.Println("-------------")
}
