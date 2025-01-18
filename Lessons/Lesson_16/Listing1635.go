func main() {
   var ord order
   var name string
   var done bool
   done = false
   fmt.Println("Welcome to Myriam's Burger Shop!")
   fmt.Print("May I have your name for the order? ")
   fmt.Scanln(&name)
   ord.name = name
   fmt.Println("Let's get your order in " + name + "!")
   for done == false {
      fmt.Println("Enter b for Burger")
      fmt.Println("Enter s for Side")
      fmt.Println("Enter d for Drink")
      fmt.Print("Enter c for Combo: ")
      choice := ""
      for contains(possibleChoices[:], choice) == false {
         fmt.Scanln(&choice)
         switch choice {
         case "b":
            fmt.Println("Burger it is!")
            var b = orderBurger()
            ord.burgers = append(ord.burgers, b)
         case "s":
            fmt.Println("Side it is!")
            var s = orderSide()
            ord.sides = append(ord.sides, s)
         case "d":
            fmt.Println("Drink it is!")
            var d = orderDrink()
            ord.drinks = append(ord.drinks, d)
         case "c":
            fmt.Println("Combo it is!")
            var c = orderCombo()
            ord.combos = append(ord.combos, c)
         default:
            fmt.Println("Unknown choice")
            fmt.Println("Please enter a valid choice")
         }
      }

      fmt.Print("Do you want to order more items? (Enter n or N to stop.): ")
      var q1 string
      fmt.Scanln(&q1)
      if strings.ToLower(q1) == "n" {
         done = true
      }
   }
   ord.display()
}
