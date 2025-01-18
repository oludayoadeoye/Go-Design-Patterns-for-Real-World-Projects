package main

import "fmt"

type account struct {
   accountNumber string
   balance float64
   owner entity
}

type entity struct{
   id string
   address string
}

// method uses value from account struct
func(acct account) HaveEnoughBalance(value float64) bool{
   if acct.balance >= value{
      return true
   }
   return false
}

func main() {
   e := entity{"000-00-0000","123 Main Street"}
   a := account{}
   a.accountNumber = "C21345345345355"
   a.balance = 140609.09
   a.owner = e

   // check if the account has 150 dollars to withdraw
   fmt.Println(a.HaveEnoughBalance(150))
   fmt.Println(a)
}