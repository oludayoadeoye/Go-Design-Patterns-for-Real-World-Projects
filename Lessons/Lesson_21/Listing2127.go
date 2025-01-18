func (dataset *Dataset) tokenize() (bool, error) {
   if dataset.empty() {
   return true, errors.New("Dataset is empty. Please read data from json first.")
   }

   for i := range dataset.reviews {
      dataset.reviews[i].Tokens = tokenize(dataset.reviews[i].ReviewText)
   }
   return false, nil
}