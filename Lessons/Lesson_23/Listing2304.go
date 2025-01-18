func returnAllAccounts(w http.ResponseWriter, r *http.Request){
   json.NewEncoder(w).Encode(Accounts) 
}