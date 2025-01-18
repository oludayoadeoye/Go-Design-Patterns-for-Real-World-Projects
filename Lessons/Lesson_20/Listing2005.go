package main

import (
   "fmt"
   "os"
)

func main() {
   f, err := os.Open("flatland01.txt")
   if err != nil {
      fmt.Println(err)
   }

   // Close the file when program is done
   defer f.Close()

   // skip the first 100 bytes
   s, err := f.Seek(100, 0)

   if err != nil { // if there is an error, print it
      fmt.Println(err) 
   }

   // display the offset
   fmt.Println(s)

   // read 20 bytes starting from the offset
   data := make([]byte, 20)
   n, err := f.Read(data)

   if err != nil { // if no error, then err is nil
      fmt.Println(err) // if there is an error, print it.
   }

   fmt.Println("Bytes read", n)
   fmt.Println("Reading starting from byte >>", s, ":", string(data[:n]))
}