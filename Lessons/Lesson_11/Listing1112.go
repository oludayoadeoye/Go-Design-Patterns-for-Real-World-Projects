package main

import (
   "fmt"
)

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
   e1 := entity{"000-00-0000", "123 Main Street"}
   fmt.Println(e1)

   e2 := entity{"000-00-0000", "123 Main Street"}
   fmt.Println(e2)

   fmt.Println(e1 == e2)

   e3 :=entity{"000-00-0000", "124 Main Street"}
   fmt.Println(e3)

   fmt.Println(e1 == e3)
}