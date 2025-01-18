package main

import (
   "fmt"
   "time"
)

func main() {

   // Declaring time in UTC
   myTime := time.Date(2025, 5, 15, 12, 0, 0, 0, time.Local)
   date1 := myTime.Add(time.Second * 6)
   date2 := myTime.Add(time.Minute * 6)
   date3 := myTime.Add(time.Hour * 6)
   date4 := myTime.Add(time.Hour * 24 * 7) 
 
   // Print the date/times output
   fmt.Println(myTime)
   fmt.Println(date1)
   fmt.Println(date2)
   fmt.Println(date3)
   fmt.Println(date4)
}
