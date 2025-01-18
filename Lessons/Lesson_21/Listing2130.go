func (dataset *Dataset) count_words() (bool, error) {
   if dataset.empty() {
      return true, errors.New("Dataset is empty. Please read data from json first.")
   }

   for i := range dataset.reviews {
      if len(dataset.reviews[i].Tokens) == 0 {
         dataset.reviews[i].Tokens = tokenize(dataset.reviews[i].ReviewText)
      }
      dataset.reviews[i].WordCount = count_words(dataset.reviews[i].Tokens)
   }
   return false, nil
}