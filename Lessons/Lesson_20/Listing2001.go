package main

import (
   "fmt"
   "io/ioutil"
)

func main() {
   // use the ReadFile from ioutil package to read the
   // entire file in memory

   data, err := ioutil.ReadFile("flatland01.txt")

   // feedback message in case of error
   if err != nil {
      fmt.Println(err)
   }

   // convert the file contents to a string and display them
   fmt.Print(string(data))
}