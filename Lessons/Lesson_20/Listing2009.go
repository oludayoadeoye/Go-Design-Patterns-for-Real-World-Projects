package main

import (
   "fmt"
   "os"
)

func main() {
   data := "Hello, world!"
   fmt.Println("data string:", data)

   // create a new file
   f, err := os.Create("another_file.txt")

   // display the new file
   if err != nil {
      panic("cannot create file: " + err.Error())
   }

   // Close the file when program is done
   defer f.Close()

   fmt.Println("new file:", f)

   // write the string to the new file
   n, err := f.WriteString(data)
   if err != nil {
      panic("cannot write to file: " + err.Error())
   }
   fmt.Println("characters in file:", n)
}
