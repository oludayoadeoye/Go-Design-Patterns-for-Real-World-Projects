func tokenize(text string) []string {
   // Set up the regexp to use a punctuation list
   re := regexp.MustCompile(`[[:punct:]]`)

   // use the ReplaceAllString to replace any punctuation with a space
   w := re.ReplaceAllString(text, " ")

   w = strings.ToLower(w) // convert to lowercase

   // Use the Fields function from the strings package to split 
   // the string w around each instance of one or more consecutive 
   // whitespace characters
   tokens := strings.Fields(w)

   // return the slice, which represents the list of tokens in 
   // order from the input string
   return tokens 
}