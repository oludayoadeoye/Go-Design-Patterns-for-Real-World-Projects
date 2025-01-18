package main

import (
   "os"
   "io/ioutil"
   "fmt"
)

func main() {
   // delete directory and all contents
   os.RemoveAll("./test_directory")

   // MkdirAll creates a tree of directories
   if err := os.MkdirAll("./test_directory/another_directory/third_directory",0755); err != nil {
      panic(err) // if there is an error then panic to fail program
   }

   // change the working directory
   if err := os.Chdir("./test_directory/another_directory/third_directory"); err != nil {
      panic(err) // if there is an error then panic to fail program
   }

   // create a file in this directory
   data := []byte("Hello, world!")
   if err := ioutil.WriteFile("new_file.txt", data, 0644); err != nil {
      panic(err) // if there is an error then panic to fail program
   }

   // list content of a directory
   content, err := ioutil.ReadDir("./")
   // if there is an error then panic to fail program.
   if err != nil {
      panic(err) // if there is an error then panic to fail program
   }

   // iterate through content
   for _, item := range content {
      fmt.Println(" ", item.Name(), item.IsDir())
   }
}