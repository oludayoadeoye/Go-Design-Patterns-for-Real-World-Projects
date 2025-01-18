package main

import (
   "chat"
   "google.golang.org/grpc"
   "log"
   "golang.org/x/net/context"
)

func main() {
   var conn *grpc.ClientConn
    conn, err := grpc.Dial(":10000", grpc.WithInsecure())
   if err != nil {
      log.Fatalf("did not connect: %s", err)
   }
   defer conn.Close() // this will execute last

   c := chat.NewChatServiceClient(conn)

   response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello from Client!"})
   if err != nil {
      log.Fatalf("Error when calling SayHello: %s", err)
   }

   // display the response from the server
   log.Printf("Response from server: %s", response.Body)

}