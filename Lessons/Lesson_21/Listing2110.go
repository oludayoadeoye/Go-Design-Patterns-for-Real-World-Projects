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
}
