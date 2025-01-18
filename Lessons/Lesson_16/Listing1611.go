func (d *drink) computePrice() int {
   if _, ok := drinks[d.getSize()]; ok {
      d.price = drinks[d.getSize()]
   }
   return d.price
}