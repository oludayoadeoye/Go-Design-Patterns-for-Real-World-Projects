package main

import(
   "fmt"
   "strings"
)

// create a function that main will call
// upper function takes as input a pointer of string
func upper(x *string) {
   *x = strings.ToUpper(*x)
}

func main() {
   var message string = "hello world"
   messageptr := &message

   upper(messageptr)

   fmt.Println(message)
}