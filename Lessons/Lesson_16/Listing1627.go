func (o *order) display() {
   fmt.Println("====================================")
   fmt.Println("===========ORDER OVERVIEW===========")
   for k, b := range o.burgers {
      fmt.Printf("=====Burger %d\n", k+1)
      b.display(true)
   }
   for k, s := range o.sides {
      fmt.Printf("=====Side %d\n", k+1)
      s.display(true)
   }
   for k, d := range o.drinks {
      fmt.Printf("=====Drink %d\n", k+1)
      d.display(true)
   }
   for k, c := range o.combos {
      fmt.Printf("=====Combo %d\n", k+1)
      c.display()
   }
   fmt.Printf("=====ORDER TOTAL: $%.2f\n", o.computePrice())
   fmt.Println("====================================")
}