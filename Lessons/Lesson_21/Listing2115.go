func count_words(words []string) map[string]int {

   word_count := make(map[string]int)

   for i := range words {
      if _, ok := word_count[words[i]]; ok {
          word_count[words[i]] = word_count[words[i]] + 1
      } else {
          word_count[words[i]] = 1
      }
   }

   return word_count
}