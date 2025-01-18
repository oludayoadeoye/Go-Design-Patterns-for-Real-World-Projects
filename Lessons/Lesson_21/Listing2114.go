func main() {
   reviews := read_json_file("./Digital_Music_5.json")
   for i := range reviews {
      tokenize(reviews[i].ReviewText) // tokenize review
   }
}