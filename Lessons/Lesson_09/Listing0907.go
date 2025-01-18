package main
 
  import "fmt"

func main() {
   // declare an empty array
   var empty [6]int

   // declare an int array and initialize its values
   var numbers = [5]int {1000, 2, 3, 7, 50}

   // declare an array without the var keyword
   words := [4]string {"hi","how","are","you"} 

   fmt.Println(empty)
   fmt.Println(numbers)
   fmt.Println(words)
}