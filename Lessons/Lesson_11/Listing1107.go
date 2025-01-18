package main

import "fmt"

type account struct {
   accountNumber string
   balance float64
}

func closeAccount(CurrentAccount *account) {
   CurrentAccount.accountNumber = "CLOSED-" + CurrentAccount.accountNumber;
}

func main() {
   var a = new(account)
   a.accountNumber = "C13242524"
   a.balance = 140.78
   fmt.Println(a.accountNumber)
   closeAccount(a)
   fmt.Println(a.accountNumber)
}