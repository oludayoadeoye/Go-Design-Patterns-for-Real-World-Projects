package main

import "fmt"

const burgerPrice = 6.00

type burger struct {
   name       string
   price      int
   condiments []string
}

func (b *burger) getName() string {
   return b.name
}

func (b *burger) computePrice() int {
   b.price = burgerPrice
   return b.price
}

func (b *burger) addCondiment(condiment string) {
   b.condiments = append(b.condiments, condiment)
}

func (b *burger) display(displayPrice bool) {
   fmt.Println("Item Name: " + b.getName())
   fmt.Print("Condiments: ")
   for _, condiment := range b.condiments {
      fmt.Print(condiment + " ")
   }
   fmt.Println()
   if displayPrice == true {
      fmt.Printf("Item Price: $%d\n", b.computePrice())
   }
}

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
