package main

import (
   "fmt"
)

type account struct {
   number string
   balance float64
}

func(a *account) withdraw(value float64) (bool, error) {
   if a.balance >= value {
      a.balance = a.balance-value
      return true, nil
   }
   return false, fmt.Errorf("Withdrawal failed, because the requested amount of %0.2f is higher than balance of %0.02f. ", value,a.balance)
}

func main() {
   acct := account{}
   acct.number = "C21345345345355"
   acct.balance = 159
   out, err :=acct.withdraw(200)
   fmt.Println(out)

   if err!=nil{
      fmt.Println(err)
      return 
   }
   fmt.Println("The withdrawal occurred successfully.")
   fmt.Println("Your new balance is", acct)
}