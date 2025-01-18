package main

import (
   "fmt"
   "github.com/gorilla/mux"
   "log"
   "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "Welcome to our API for Smart Data!")
   fmt.Println("Endpoint: /")
}

func getGeoLocationData(w http.ResponseWriter, r *http.Request) {
   log.Println("Incoming API GeoLocation Request")
}

func getQuote(w http.ResponseWriter, r *http.Request) {
   log.Println("Incoming API Quote Request")
}

// handleRequests will process HTTP requests and redirect them to
// the appropriate Handle Function
func handleRequests() {
   // create a router to handle our requests from the mux package.
   router := mux.NewRouter().StrictSlash(true)
   // access root page
   router.HandleFunc("/", homePage)
   router.HandleFunc("/getGeoLocationData/{address}", getGeoLocationData)
   router.HandleFunc("/getQuote/{symbol}", getQuote)
   log.Fatal(http.ListenAndServe(":11112", router))
}

func main() {
   handleRequests()
}