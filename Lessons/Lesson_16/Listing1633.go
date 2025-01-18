func orderDrink() drink {
   fmt.Print("These are the available drinks: ")
   fmt.Println(drinkTypes)
   fmt.Print("These are the available sizes: ")
   fmt.Println("[12 16 24]")
 
   var choice bool = false
   var drinkTypeChoice string
   var drinkSizeChoice int
   for choice == false {
      fmt.Print("What drink do you want? ")
      fmt.Scanln(&drinkTypeChoice)
      if contains(drinkTypes, strings.ToUpper(drinkTypeChoice)) {
         choice = true
      } else {
         fmt.Println("Please enter a valid drink")
      }
   }
   choice = false
   for choice == false {
      fmt.Print("What size do you want? ")
      fmt.Scanln(&drinkSizeChoice)
      if _, ok := drinks[drinkSizeChoice]; ok {
         choice = true
      } else {
         fmt.Println("Please enter a valid size")
      }
   }
   var d drink
   d.name = strings.ToLower(drinkTypeChoice)
   d.size = drinkSizeChoice
   d.computePrice() // equivalent also to d.price = drinks[drinkSizeChoice]
   return d
}