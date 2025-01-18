func handleRequests() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", homePage)
    router.HandleFunc("/accounts", returnAllAccounts)
    router.HandleFunc("/account/{number}", returnAccount)
    log.Fatal(http.ListenAndServe(":10000", router))
}