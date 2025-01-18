func getGeoLocationData(w http.ResponseWriter, r *http.Request) {
   log.Println("Incoming API GeoLocation Request")
   vars := mux.Vars(r)
   address := vars["address"]
   var conn *grpc.ClientConn
   conn, err := grpc.Dial(":9997", grpc.WithInsecure())
   defer conn.Close()
   if err != nil {
      fmt.Fprintln(w, "Error. Please try again later")
      return
   }
   c := geolocation.NewGeoLocationServiceClient(conn)
   response, err := c.GetGeoLocationData(context.Background(), &geolocation.Message{Body: address})
   if err != nil {
      fmt.Fprintln(w, err)
      return
   }
   fmt.Fprintln(w, response.Body)
} 