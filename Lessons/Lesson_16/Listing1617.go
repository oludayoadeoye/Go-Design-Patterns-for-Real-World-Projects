func (s *side) display(displayPrice bool) {
   fmt.Println("Item Name: " + s.getName())
   if displayPrice == true {
      fmt.Printf("Item Price: $%d\n", s.computePrice())
   }
}