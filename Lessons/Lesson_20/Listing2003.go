package main

import (
   "fmt"
   "os"
)

func main() {
   f, err := os.Open("flatland01.txt")

   if err != nil {
      fmt.Println(err) // if there is an error, print it
   }

   // create a slice of bytes
   b1 := make([]byte, 5) 
   data, err := f.Read(b1)

   // feedback message in case of error; otherwise nil
   if err != nil {
      fmt.Println(err) // if there is an error, print it
   }

   // display the slice
   fmt.Println(string(b1[:data]))

   // close the file after completing the operations
   f.Close()
}
