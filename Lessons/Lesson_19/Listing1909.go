package main

import (
   "fmt"
   "time"
)

func main() {
   // display current time
   now := time.Now()
   fmt.Println("Current date and time:", now)

   // create custom time
   customTime := time.Date(
      2025, 05, 15, 15, 20, 00, 0, time.Local)
   fmt.Println("Custom date and time:", customTime)

   // comparisons
   fmt.Println("The current time is before the custom time:", now.Before(customTime))
   fmt.Println("The current time is after the custom time:", now.After(customTime))
   fmt.Println("The current time is equal to the custom time:", now.Equal(customTime))
}