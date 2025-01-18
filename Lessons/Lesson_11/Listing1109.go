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

func main() {
   e := entity{"000-00-0000", "123 Main Street"}
   a := account{}
   a.accountNumber ="C21345345345355"
   a.balance = 140609.09

   // assign the entity struct as a value in the account struct
   a.owner = e

   fmt.Println(a)
}