func (o *order) computePrice() int {
   var price = 0
   for _, b := range o.burgers {
      price = price + b.computePrice()
   }
   for _, s := range o.sides {
      price = price + s.computePrice()
   }
   for _, d := range o.drinks {
      price = price + d.computePrice()
   }
   for _, c := range o.combos {
      price = price + c.computePrice()
   }
   o.price = price
   return o.price
}