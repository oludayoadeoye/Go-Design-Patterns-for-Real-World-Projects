func contains(arr []string, choice string) bool {
   for _, v := range arr {
      if v == choice {
         return true
      }
   }
   return false
}