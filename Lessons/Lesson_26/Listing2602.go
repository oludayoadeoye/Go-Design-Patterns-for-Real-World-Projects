package mymodule

import "testing"

func TestRepeatString(t *testing.T) {
   expected := "AA"
   result := RepeatString("A", 2)
   if result != "AA" {
      t.Errorf("Error: Expected = %s, Result = %s", expected, result)
   }
}