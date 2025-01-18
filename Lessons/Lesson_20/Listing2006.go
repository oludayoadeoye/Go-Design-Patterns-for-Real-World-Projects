package main

import (
   "bufio"
   "fmt"
   "os"
)

func main() {
   f, err := os.Open("flatland01.txt")
   if err != nil {
      panic(err)
   }

   br := bufio.NewReader(f) // create a buffered reader
   data, err2 := br.Peek(5) // read 5 bytes

   if err2 != nil {
      fmt.Println(err)
   }

   // display the peeked data
   fmt.Println(string(data))
}