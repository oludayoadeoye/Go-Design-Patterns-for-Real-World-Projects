func (c *combo) display() {
   fmt.Println("Burger For Combo")
   c.burger.display(false)
   fmt.Println("Side For Combo")
   c.side.display(false)
   fmt.Println("Drink For Combo")
   c.drink.display(false)
   fmt.Printf("Price For Combo: $%d\n", c.computePrice())
}