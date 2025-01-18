package main

import "fmt"

type account struct {
   accountNumber string
   balance float64
}

func main() {
   a := account{ balance: 140.78, accountNumber: "C13242524"}

   b := account{ accountNumber: "S12212321"}

   fmt.Println(a)
   fmt.Println(b)
}