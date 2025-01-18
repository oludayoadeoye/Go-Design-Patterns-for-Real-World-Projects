package mymodule

func RepeatString(text string, count int) string {
   if count < 2 { //even if user enters negative value or 0, we return the text by default
      return text
   }
   out := ""
   for i := 0; i < count; i++ {
      out = out + text
   }
   return out
}