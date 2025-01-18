package main

import (
   "os"
)

func main() {
   // create a directory
   err := os.Mkdir("./test_directory", 0755) // will throw an error if the directory exists

   // if there is an error then panic to fail program
   if err != nil {
      panic(err)
   }
}