package main

import "fmt"

type AccountOperations interface{
   // methods
   withdraw(value float64) bool
   deposit(value float64) bool
   displayInfo()
}

type account struct {
   number string
   balance float64
}

func (acct *account) withdraw(value float64) bool {
   if acct.balance >= value{
      acct.balance = acct.balance - value
      return true
   }
   return false
}

func (acct *account) deposit(value float64) bool {
   if value > 0 {
      acct.balance = acct.balance + value
      return true
   }
   return false
}

func (acct *account) displayInfo() {
   fmt.Println("Account Balance:", acct.balance)
   fmt.Println("Account Number: ", acct.number)
}

func main() {
   var ao AccountOperations
   fmt.Println("initial value:", ao) 

   // Assign a pointer to an account value that is created to ao 
   // We can only do this because the account type implements 
   // all the methods of AccountOperations Interface

   ao = &account{"C13443533535",1500}

   //withdrawal amount
   ao.withdraw(150)
   ao.displayInfo()

   // deposit amount
   ao.deposit(1000)
   ao.displayInfo()
}