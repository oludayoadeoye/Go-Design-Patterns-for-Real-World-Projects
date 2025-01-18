package main

import (
   "os"
   "fmt"
)

func main() {
   data, err := ioutil.ReadFile("badFileName.txt")
   if err != nil {
      panic(err)
   }

   // Won't get here if there is an error reading file
   fmt.Print(string(data))
}