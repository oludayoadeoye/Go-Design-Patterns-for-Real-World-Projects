func getQuote(w http.ResponseWriter, r *http.Request) {
   log.Println("Incoming API Quote Request")
   vars := mux.Vars(r)
   symbol := vars["symbol"]
   var conn *grpc.ClientConn
   conn, err := grpc.Dial(":9997", grpc.WithInsecure())
   defer conn.Close()
   if err != nil {
      // issues with connecting to gRPC
      fmt.Fprintln(w, "Error. Please try again later") //provide standard 
      // error messages instead of technical errors
      return
   }
   c := finance.NewFinanceServiceClient(conn)
   response, err := c.GetQuoteData(context.Background(), &finance.Message{Body: symbol})
   if err != nil {
      // issues with the symbol (symbol doesn't exist) or the Yahoo 
      // API is not working.
      fmt.Fprintln(w, err)
      return
   }
   fmt.Fprintln(w, response.Body) //return the quote in the body of the response
}