package main

import (
   "./finance/finance"
   "./geolocation/geolocation"
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
   // adding geolocation server
   ch1 := geolocation.Server{}
   geolocation.RegisterGeoLocationServiceServer(grpcServer, &ch1)
   // adding finance server
   ch2 := finance.Server{}
   finance.RegisterFinanceServiceServer(grpcServer, &ch2)
   err = grpcServer.Serve(listener)
   if err != nil {
      log.Fatalf("failed to serve: %s", err)
   }
}