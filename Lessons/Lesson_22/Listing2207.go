package main

import (
   "fmt"
)

func FormatAmount(a float64) string {
   // use %.2f for precision 2 which is adequate to represent
   // dollar amounts for now
   return "USD " + fmt.Sprintf("%.2f", a)
}

func main(){
   fmt.Println("main function")
   fmt.Println(FormatAmount(2.00))
}