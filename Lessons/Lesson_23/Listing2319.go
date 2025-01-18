package main

import (
   "encoding/json"
   "fmt"
   "log"
   "net/http"
   "github.com/gorilla/mux"
   "io/ioutil"
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
// use the ResponseWriter w to display some text when we visit the home page
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
   // create new account
   router.HandleFunc("/account", createAccount).Methods("POST")
   // delete requested account
   router.HandleFunc("/account/{number}", deleteAccount).Methods("DELETE")
   // define the localhost
   log.Fatal(http.ListenAndServe(":10000", router))
}

// return the dataset Accounts in a JSON format
func returnAllAccounts(w http.ResponseWriter, r *http.Request){
   // we use the Encode function to convert the Account slice into a json object
   json.NewEncoder(w).Encode(Accounts)
}

// return a single account
func returnAccount(w http.ResponseWriter, r *http.Request){
   vars := mux.Vars(r)
   key := vars["number"]
   for _, account := range Accounts {
      if account.Number == key {
         json.NewEncoder(w).Encode(account)
      }
   }
}

// create a new account
func createAccount(w http.ResponseWriter, r *http.Request) {
   reqBody, _ := ioutil.ReadAll(r.Body)
   var account Account
   json.Unmarshal(reqBody, &account)
   Accounts = append(Accounts, account)
   json.NewEncoder(w).Encode(account)
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
   // use mux to parse the path parameters
   vars := mux.Vars(r)
   // extract the account number of the account we wish to delete
   id := vars["number"]
   // we then need to loop through dataset
   for index, account := range Accounts {
      // if our id path parameter matches one of our
      // account numbers
      if account.Number == id {
         // updates our dataset to remove the account
         Accounts = append(Accounts[:index], Accounts[index + 1:]...)
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