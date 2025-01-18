package main

import (
   "fmt"
   "regexp"
)

func main() {
   text := "Hello, Sean! -How are you?"
   fmt.Println("original string: " + text)

   // The following is regex for the punctuation list. This
   // means that any of the punctuation in the list will be
   // replaced by a space
   re := regexp.MustCompile(`[.,!?\-_#^()+=;/&' ̋]`)

   // Use the ReplaceAllString to replace any punctuation with a space
   w := re.ReplaceAllString(text, " ")

   fmt.Println("string after replacing punctuation with a space: " + w)
}