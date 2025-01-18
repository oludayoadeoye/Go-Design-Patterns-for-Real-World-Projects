package main

import "fmt"

func main() {
   var s interface{}
   fmt.Println(s)
   fmt.Printf("s is nil and has type %T value %v\n", s, s)
} 