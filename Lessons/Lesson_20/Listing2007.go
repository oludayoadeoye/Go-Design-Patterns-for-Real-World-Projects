package main

import (
   "bufio"
   "fmt"
   "os"
)

func main() {
   file, err := os.Open("flatland01.txt")
   if err != nil {
      // if error is not nil, panic
      panic("File not found")
   }

   // Close the file when program is done 
   defer file.Close()

   // create a scanner to read from file and split text based on lines
   scanner := bufio.NewScanner(file)
   scanner.Split(bufio.ScanLines)

   // create a slice, which will contain the lines read from file
   var lines []string

   // use Scan to iterate through the file
   for scanner.Scan() {
      // append the current line to the slice lines
      lines = append(lines, scanner.Text())
   }
   // iterate through lines
   for _, line := range lines {
      fmt.Println("line:", line)
   }
}