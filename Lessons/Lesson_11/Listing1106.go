package main

import "fmt"

type account struct {
   accountNumber string
   balance float64
}

func main() {
   a := new(account)
   a.accountNumber = "C14235345354"
   a.balance = 140

   fmt.Println(a)
}