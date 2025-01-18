package main

import (
   "os"
)

func main() {
   // delete directory and all contents
   os.RemoveAll("./test_directory")

   // create a directory
   err := os.Mkdir("./test_directory", 0755)

   // if there is an error then panic to fail program
   if err != nil {
      panic(err)
   }

   // MkdirAll creates a tree of directories
   err = os.MkdirAll("./test_directory/another_directory/third_directory", 0755)
   // if there is an error then panic to fail program
   if err != nil {
      panic(err)
   }
}