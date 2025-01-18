package finance

import (
   "golang.org/x/net/context"
   "log"
)

type Server struct {
}

func (s *Server) GetQuoteData(ctx context.Context, in *Message) (*Message, error) {
   log.Println("Incoming Quote Request")
   return &Message{Body: "Hello"}, nil
}