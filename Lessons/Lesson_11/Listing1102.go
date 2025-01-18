package main

import "fmt"

type account struct {
   accountNumber string
   balance float64
}

func main() {
   var a account

   a.balance = 140
   a.accountNumber = "C14235345354"

   fmt.Println(a)

   fmt.Println("The account number is", a.accountNumber)

   fmt.Println("The current balance is", a.balance)
}