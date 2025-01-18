package main

import (
   "fmt"
   "strconv"
)

func main() {
   var str string = "10x"
   // the ParseInt function returns the parsed integer or
   // the error if the conversion failed
   nbr, error := strconv.ParseInt(str,10,8)
   fmt.Println(nbr)
   fmt.Println(error)
}