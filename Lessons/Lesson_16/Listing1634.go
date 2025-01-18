func orderCombo() combo {
   var c combo
   fmt.Println("Let's get you a combo meal!")
   fmt.Println("First, let's order the burger for your combo")
   c.burger = orderBurger()

   fmt.Println("Now, let's order the drink for your combo")
   c.drink = orderDrink()

   fmt.Println("Finally, let's order the side for your combo")
   c.side = orderSide()

   return c
}