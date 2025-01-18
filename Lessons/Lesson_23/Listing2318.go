func handleRequests() {
   router := mux.NewRouter().StrictSlash(true)
   router.HandleFunc("/", homePage)
   router.HandleFunc("/accounts", returnAllAccounts)
   router.HandleFunc("/account/{number}", returnAccount)
   router.HandleFunc("/account", createAccount).Methods("POST")
   router.HandleFunc("/account/{number}", deleteAccount).Methods("DELETE")
   log.Fatal(http.ListenAndServe(":10000", router))
}