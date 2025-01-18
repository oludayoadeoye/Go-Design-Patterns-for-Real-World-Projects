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