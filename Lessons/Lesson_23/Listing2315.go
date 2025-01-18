func handleRequests() {
   router := mux.NewRouter().StrictSlash(true)
   router.HandleFunc("/", homePage)
   router.HandleFunc("/accounts", returnAllAccounts)
   router.HandleFunc("/account/{number}", returnAccount)
   router.HandleFunc("/account", createAccount).Methods("POST")
   // our API will be accessible from http://localhost:10000/
   // we add the router as a handler in the ListenAndServe function
   log.Fatal(http.ListenAndServe(":10000", router))
}