package main

import (
   "fmt"
   "time"
)

func main() {
   // display current time
   now := time.Now()
   fmt.Println("Current date and time:", now)

   newDate := now.AddDate(0, 0, 14)
   fmt.Println("Two weeks in the future:", newDate)

   // Using a custom date
   customDate := time.Date(
   2025, 05, 15, 15, 20, 00, 0, time.Local)
   fmt.Println("Custom date and time:", customDate)

   newCustomDate := customDate.AddDate(0, 0, 14)
   fmt.Println("Two weeks after custom date:", newCustomDate)
}