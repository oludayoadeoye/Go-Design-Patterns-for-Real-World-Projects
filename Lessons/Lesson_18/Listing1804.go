package main

import (
   "fmt"
   "time"
)

func goroutine(limit int) {
   for i := 0;i < limit; i++ {
      time.Sleep(250 * time.Millisecond)
      fmt.Print(i)
      fmt.Println(" - calling goroutine")
   }
}

func anothergoroutine(limit int) {
   for i := 0;i < limit; i++ {
      time.Sleep(400 * time.Millisecond)
      fmt.Print(i)
      fmt.Println(" - calling anothergoroutine")
   }
 }

func main() {
   fmt.Println("Starting...")
   go goroutine(10)
   go anothergoroutine(10)

   time.Sleep(6 * time.Second)
   fmt.Println("main goroutine")
}