package main

import "fmt"

type account struct {
   accountNumber string
   balance float64
}

func main() {
   a := account{"C13242524", 140.78}
   fmt.Println(a)

   b := account{}
   fmt.Println(b)
}
