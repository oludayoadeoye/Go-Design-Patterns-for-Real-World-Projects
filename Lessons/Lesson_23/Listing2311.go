func returnAccount(w http.ResponseWriter, r *http.Request){
   vars := mux.Vars(r)
   key := vars["number"]
   for _, account := range Accounts {
      if account.Number == key {
         json.NewEncoder(w).Encode(account)
      }   
   }
}