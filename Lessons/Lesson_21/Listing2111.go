package main

import (
   "fmt"
   "regexp"
   "strings"
)

func main() {
   text := "Hello, Sean! -How are you?"
   fmt.Println("original string: " + text)

   re := regexp.MustCompile(`[[:punct:]]`) 

   w := re.ReplaceAllString(text, " ")

   fmt.Println("string after replacing punctuation with a space: " + w)

   w = strings.ToLower(w) // convert to lowercase
   // Use the Fields function from the strings package to split 
   // the string w around each instance of one or more consecutive 
   // whitespace characters
   tokens := strings.Fields(w) 

   fmt.Print("Tokens: ")
   fmt.Println(tokens)
}
