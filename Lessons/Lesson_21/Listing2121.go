func (dataset *Dataset) tokenize() {
   for i := range dataset.reviews {
      dataset.reviews[i].Tokens = tokenize(dataset.reviews[i].ReviewText)
   }
}