package main

import "fmt"

type account struct {
   accountNumber string
   balance float64
}

func main() {
   a := account{"C21345345345355", 15470.09}
   p := &a

   (*p).balance = 220 
   fmt.Println(a)

   p.balance = 320 
   fmt.Println(a)
}