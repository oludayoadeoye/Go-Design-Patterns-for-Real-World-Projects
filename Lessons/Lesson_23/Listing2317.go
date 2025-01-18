func deleteAccount(w http.ResponseWriter, r *http.Request) {
   // use mux to parse the path parameters
   vars := mux.Vars(r)
   // extract the account number of the account we wish to delete
   id := vars["number"]
   // we then need to loop through the dataset
   for index, account := range Accounts {
      // if our id path parameter matches one of our
      // account numbers
      if account.Number == id {
         // updates our dataset to remove the account
         Accounts = append(Accounts[:index], Accounts[index + 1:]...)
      }
   }
}