package main

import (
   "fmt"
   "io/ioutil"
)

func main() {
   // create a slice of bytes (UTF-­8 code) from an input string
   data := []byte("Hello, world!")

   fmt.Println(data) // display the slice

   // feedback message in case of error
   err := ioutil.WriteFile("new_file.txt", data, 0644)
   if err != nil {
      panic("cannot write file: " + err.Error())
   }

   new_file, err := ioutil.ReadFile("new_file.txt")
   // feedback message in case of error
   if err != nil {
      panic("cannot read file: " + err.Error())
   }

   // convert the file contents to a string and display them
   fmt.Print(string(new_file))
}