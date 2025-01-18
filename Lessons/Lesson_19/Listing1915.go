package main

import (
   "fmt"
   "time"
)

func main() {
   now := time.Now()
   unixtime := now.Unix() // Unix time
   unixnanotime := now.UnixNano() // Unix nano time
   fmt.Println(now)
   fmt.Println(unixtime)
   fmt.Println(unixnanotime)
}