package main

import "fmt"

func main() {
   matrix := [3][3]int {
      {1, 2, 3},
      {4, 5, 6},
      {7, 8, 9},
   }

   matrix_2 := [3][3]int {
      {1, 2, 3},
      {4, 5, 6},
      {7, 8, 9},
   }

   matrix_3 := [3][3]int { 
      {9, 9, 9},
      {9, 9, 9},
      {9, 9, 9},
   }

   // Compare matrix to matrix_1
   if( matrix == matrix_2 ) {
      fmt.Println("matrix equals matrix_2")
   } else {
      fmt.Println("matrix does NOT equal matrix_2")
   }

   // Compare matrix to matrix_3
   if( matrix == matrix_3 ) {
      fmt.Println("matrix equals matrix_3")
   } else {
      fmt.Println("matrix does NOT equal matrix_3")
   }

   // Print out the three arrays
   fmt.Println(matrix)
   fmt.Println(matrix_2)
   fmt.Println(matrix_3)
}