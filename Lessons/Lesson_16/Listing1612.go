func (d *drink) display(displayPrice bool) {
   fmt.Println("Item Name: " + strings.ToUpper(d.getName()))
   fmt.Printf("Item Size: %d\n", d.size)
   if displayPrice == true {
      fmt.Printf("Item Price: $%d\n", d.computePrice())
   }
}