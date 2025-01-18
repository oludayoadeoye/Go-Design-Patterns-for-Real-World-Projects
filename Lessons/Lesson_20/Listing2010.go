package main

import (
   "bufio"
   "fmt"
   "os"
   "io/ioutil"
)

func main() {
   data :="Hello, world!!!"
   fmt.Println("original string:", data)

   // create a new file
   f, err := os.Create("another_file.txt")
   if err != nil {
      panic("cannot create file: " + err.Error())
   }

   // Close the file when program is done
   defer f.Close()

   // create a buffered writer that we can use to write data to the new file
   bw := bufio.NewWriter(f)

   // write the data to the buffered writer
   n, err := bw.WriteString(data)
   if err != nil {
      panic("cannot write string: " + err.Error())
   }

   // display the number of bytes written
   fmt.Println("bytes written:", n)

   // flush flushes/submits the data to the underlying io.Writer
   bw.Flush()

   newFile, err := ioutil.ReadFile("another_file.txt")

   // feedback message in case of error
   if err != nil {
      panic("cannot read file: " + err.Error())
   }

   // convert the file contents to a string and display them
   fmt.Print("file contents: ", string(newFile))
}