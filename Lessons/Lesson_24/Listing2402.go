package main

import (
   "fmt"
   "google.golang.org/grpc" // import the grpc package
   "log"
   "net"
)

func main() { 
   fmt.Println("Our first gRPC server ")
   listener, err := net.Listen("tcp", ":10000")

   if err != nil {
      log.Fatalf("failed to listen: %v", err)
   }

   // create a new grpc server
   grpcServer := grpc.NewServer()
   err = grpcServer.Serve(listener)
   if err != nil {
      log.Fatalf("failed to serve: %s", err)
   }
}