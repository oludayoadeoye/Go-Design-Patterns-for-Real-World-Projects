func (dataset *Dataset) count_words() {
   for i := range dataset.reviews {
      dataset.reviews[i].WordCount = count_words(dataset.reviews[i].Tokens)
   }
}