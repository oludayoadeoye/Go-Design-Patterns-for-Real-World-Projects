package main

import (
   "encoding/json"
   "fmt"
   "log"
   "net/http"
   "github.com/gorilla/mux"
)

// we create a type Account that will be used to represent a bank account
type Account struct {
   Number string `json:"AccountNumber"`
   Balance string `json:"Balance"`
   Desc string `json:"AccountDescription"`
}

// we use Accounts as a global variable because it is used by
// several functions in the code
var Accounts []Account

// implement the homePage
// use the ResponseWriter w to display some text when we visit the homepage
func homePage(w http.ResponseWriter, r *http.Request){
   fmt.Fprintf(w, "Welcome to our bank!")
   // we can use a print command to log the request or we can log it to a file, etc.
   fmt.Println("Endpoint: /")
}

// handleRequests will process HTTP requests and redirect them to
// the appropriate Handle function
func handleRequests() {
   // create a router to handle our requests from the mux package.
   router := mux.NewRouter().StrictSlash(true)
   // access root page
   router.HandleFunc("/", homePage)
   // returnAllAccounts
   router.HandleFunc("/accounts", returnAllAccounts)
   // return requested account
   router.HandleFunc("/account/{number}", returnAccount)
   // define the localhost
   log.Fatal(http.ListenAndServe(":10000", router))
}

// return the dataset Accounts in a JSON format
func returnAllAccounts(w http.ResponseWriter, r *http.Request){
   // we use the Encode function to convert the Account slice into a json object
   json.NewEncoder(w).Encode(Accounts)
}

func returnAccount(w http.ResponseWriter, r *http.Request){
   vars := mux.Vars(r)
   key := vars["number"]
   for _, account := range Accounts {
      if account.Number == key {
         json.NewEncoder(w).Encode(account)
      }
   }
}

func main() {
   // initialize the dataset
   Accounts = []Account{
      Account{Number: "C45t34534", Balance: "24545.5", Desc: "Checking Account"},
      Account{Number: "S3r53455345", Balance: "444.4", Desc: "Savings Account"},
   }

   // execute handleRequests, which will kick off the API
   // we can access the API using the the URL defined above
   handleRequests()
}