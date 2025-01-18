package main

import "fmt"

// create first interface
type AccountOperations interface{
   computeInterest() float64
   displayInfo()
}

// create second interface
type UserOperations interface{
   changeANumber(number string)
}

// create a struct type
type SavingsAccount struct {
   number string
   balance float64
   interest float64
}

// implement method from first interface 
func(a *SavingsAccount) computeInterest() float64{
   return 0.005
}

// implement method from first interface 
func(a *SavingsAccount) displayInfo() {
   fmt.Println(a.number)
   fmt.Println(a.balance)
   fmt.Println(a.interest)
}

// implement method from second interface 
func(a *SavingsAccount) changeANumber(number string) {
   a.number=number
}

func main() {
   // create a SavingsAccount variable
   acct := SavingsAccount{}
   acct.number = "S21345345345355"
   acct.balance = 159

   // Declare an interface variable for AccountOperations
   var ao1 AccountOperations

   // acct implements the method of interface AccountOperations

   ao1 = &acct
   fmt.Println("ao1 info:")
   ao1.displayInfo()

   fmt.Println("--------------") // print divider for output

   // Declare an interface variable for UserOperations
   var uo1 UserOperations

   // acct also implements the methods of interface UserOperations
   uo1 = &acct
   // execute the account number change
   uo1.changeANumber("2345353453")

   fmt.Println("updated ao1 info:")
   ao1.displayInfo() 
}