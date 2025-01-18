package main

import (
   "fmt"
)

func FormatAmount(a float64) string {
   // handle when a is negative
   if a < 0 {
       return "Impossible operation"
   }
   // use %.2f for precision 2 which is enough to represent dollar amounts for now
   return "USD " + fmt.Sprintf("%.2f", a)
}

func SubtractFormatAmount(a, b float64) string {
   // if a is negative
   if a < 0 {
       return "Impossible operation"
   }
   // if b is negative
   if b < 0 {
       return "Impossible operation"
   }
   if a >= b {
      return "USD " + fmt.Sprintf("%.2f", a - b)
   }
   // if 0 < a < b
   return "Impossible operation"
}

func main(){
   fmt.Println(FormatAmount(0.00))
   fmt.Println(FormatAmount(-10.00))
   fmt.Println(SubtractFormatAmount(0.03 , 0.42))
   fmt.Println(SubtractFormatAmount(-0.03 , -0.42))
}