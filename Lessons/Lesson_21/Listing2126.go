func (dataset *Dataset) empty() bool {
   if len(dataset.reviews) == 0 {
      return true
   }
   return false
}