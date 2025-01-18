func createAccount(w http.ResponseWriter, r *http.Request) {
   reqBody, _ := ioutil.ReadAll(r.Body)
   var account Account
   json.Unmarshal(reqBody, &account)
   Accounts = append(Accounts, account)
   json.NewEncoder(w).Encode(account)
}