func handleRequests() {
   // create a router to handle our requests from the mux package.
   router := mux.NewRouter().StrictSlash(true)
   router.HandleFunc("/", homePage) 
   router.HandleFunc("/accounts", returnAllAccounts) 
   log.Fatal(http.ListenAndServe(":10000", router)) 
}