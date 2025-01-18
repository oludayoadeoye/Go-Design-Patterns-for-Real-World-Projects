package main

import "fmt"

func main() {
   var myChannel = make(chan int)
   fmt.Printf("Channel Type is %T", myChannel)
}