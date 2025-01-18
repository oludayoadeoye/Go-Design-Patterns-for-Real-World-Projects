package main

import (
   "fmt"
   "google.golang.org/grpc"
   "log"
   "net"
)

func main() {
   fmt.Println("Smart Data Server ")
   listener, err := net.Listen("tcp", ":9997")
   if err != nil {
      log.Fatalf("failed to listen: %v", err)
   }
   grpcServer := grpc.NewServer()
   err = grpcServer.Serve(listener)
   if err != nil {
      log.Fatalf("failed to serve: %s", err)
   }
}  