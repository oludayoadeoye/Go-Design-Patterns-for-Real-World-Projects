package main

import "fmt"

// create the struct
type account struct {
   accountNumber string
   balance float64
}

func main() {
   var a account

   a.balance = 140
   a.accountNumber = "C14235345354"

   fmt.Println(a)
}