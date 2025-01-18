package main

import (
   "fmt"
   "time"
)

func goroutine() {
   fmt.Println("This is my first goroutine.")
}

func anothergoroutine() {
   fmt.Println("This is my second goroutine.")
}

func main() {
   fmt.Println("Starting...")

   go goroutine()
   go anothergoroutine()

   time.Sleep(4* time.Second)
   fmt.Println("main Goroutine")
}