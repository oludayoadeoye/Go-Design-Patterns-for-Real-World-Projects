func handleRequests() {
   http.HandleFunc("/", homePage) 
   http.HandleFunc("/accounts", returnAllAccounts) 
   log.Fatal(http.ListenAndServe(":10000", nil))
}