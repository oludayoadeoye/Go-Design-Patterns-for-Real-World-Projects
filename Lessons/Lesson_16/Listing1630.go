func orderBurger() burger {
   var b burger
   b.name = "Beef Burger"
   fmt.Print("Do you want condiments on your burger? (type y for yes): ")
   var choice1 string
   fmt.Scanln(&choice1)
   if strings.ToLower(choice1) == "y" {
      for _, condiment := range burgerCondiments {
         var choice2 string
         fmt.Print("Do you want " + condiment + " on your burger? (type y for yes): ")
         fmt.Scanln(&choice2)
         if strings.ToLower(choice2) == "y" {
            b.addCondiment(condiment)
         }
      }
   }
   return b
}