package main

import "fmt"

type AccountOperations interface{
   // Methods
   computeInterest() float64
}

type SavingsAccount struct {
   number string
   balance float64
}

type CheckingAccount struct {
   number string
   balance float64
}

func(acct *SavingsAccount) computeInterest() float64{
   return 0.005
}

func(acct *CheckingAccount) computeInterest() float64{
   return 0.001
}

func main() {
   acct := SavingsAccount{}
   acct.number="S21345345345355"
   acct.balance=159

   var ao1 AccountOperations
   ao1 = &acct
   fmt.Println("savings interest:", ao1.computeInterest())

   acct2 := CheckingAccount{}
   acct2.number = "C218678678345345355"
   acct2.balance = 2000

   var ao2 AccountOperations
   ao2 = &acct2
   fmt.Println("checking interest:", ao2.computeInterest())
}