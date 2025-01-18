package main

import "fmt"

type account struct {
   accountNumber string
   balance float64
}

func main() {
   var a = account{"C14235345354",140}
   fmt.Println(a)
}