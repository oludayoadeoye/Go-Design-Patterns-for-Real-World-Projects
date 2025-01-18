func main() {
   var b burger
   b.name = "Burger"
   b.addCondiment("Lettuce")
   b.addCondiment("Tomato")
   b.addCondiment("Onion")
   b.addCondiment("Mayo")
   b.computePrice()
   b.display(true)
}
