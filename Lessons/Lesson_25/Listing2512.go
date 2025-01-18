var YAHOO_API_KEY string = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
var URL = "https://yfapi.net/v6/finance/quote?region=US&lang=en&symbols="

func BuildYahooRequest(symbol string) (*http.Response, error) {
   req, err := http.NewRequest("GET", URL+symbol, nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Accept", "application/json")
   req.Header.Set("X-Api-Key", YAHOO_API_KEY)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   return resp, nil
}