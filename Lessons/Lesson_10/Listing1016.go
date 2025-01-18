package main

import "fmt"

func main() {
   fmt.Print("Enter your name: ") 

   var name string
   fmt.Scanln(&name) 

   fmt.Print("Enter your age: ")
   var age string
   fmt.Scanln(&age)

   fmt.Print("Enter your gender: ")
   var gender string
   fmt.Scanln(&gender)

   fmt.Print("Your first name is: ")
   fmt.Println(name)

   fmt.Print("Your age is: ")
   fmt.Println(age)

   fmt.Print("Your gender is: ")
   fmt.Println(gender)
}