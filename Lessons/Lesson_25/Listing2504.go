package geolocation

import (
   "golang.org/x/net/context"
   "log"
)

type Server struct {
}

func (s *Server) GetGeoLocationData(ctx context.Context, in *Message) 
(*Message, error) {
   log.Println("Incoming GeoLocation Request")
   return &Message{Body: "Hello"}, nil
}
  