package main

import (
   "fmt"
   "log"
   "net"
)

func main() {
   listener, err := net.Listen("tcp", ":10000") 
   fmt.Println(listener)

   if err != nil {   
      log.Fatalf("failed to listen: %v", err)
   }
}