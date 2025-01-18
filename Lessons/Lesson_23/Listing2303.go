func homePage(w http.ResponseWriter, r *http.Request){
   fmt.Fprintf(w, "Welcome to our bank!")
   fmt.Println("Endpoint: /") 
}