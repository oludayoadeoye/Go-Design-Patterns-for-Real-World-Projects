func (c *combo) computePrice() int {
   c.price = c.burger.computePrice() + c.drink.computePrice() + c.side.computePrice() - comboDiscount
   return c.price
}