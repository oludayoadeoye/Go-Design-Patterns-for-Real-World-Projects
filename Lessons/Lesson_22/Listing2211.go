package main

import (
   "fmt"
)

func FormatAmount(a float64) string {
   // use %.2f for precision 2 which is adeqiate to represent
   // dollar amounts for now
   return "USD " + fmt.Sprintf("%.2f", a)
}

func SubtractFormatAmount(a, b float64) string {
   return "USD " + fmt.Sprintf("%.2f", a - b)
}

func main(){
   fmt.Println("main function")
   fmt.Println(FormatAmount(2.00))
   fmt.Println(SubtractFormatAmount(2.00, 1.14))
}