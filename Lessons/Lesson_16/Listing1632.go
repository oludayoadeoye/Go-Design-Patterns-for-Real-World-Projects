func orderSide() side {
   fmt.Print("These are the available sides: ")
   fmt.Println(sideTypes)
   var choice bool = false
   var sideTypeChoice string
   for choice == false {
      fmt.Print("What side do you want? ")
      fmt.Scanln(&sideTypeChoice)
      if contains(sideTypes[:], sideTypeChoice) {
         choice = true
      } else {
         fmt.Println("Please enter a valid choice")
      }
   }
   var s side
   s.name = strings.ToLower(sideTypeChoice)
   s.computePrice()
   return s
}
