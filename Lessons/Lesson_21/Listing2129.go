func main() {
   dataset := Dataset{filepath: "./Digital_Music_5.json"}
   _, err := dataset.read_json_file()
   if err != nil {
      log.Fatal(err.Error())
   }

   //dataset.tokenize()
   dataset.count_words()
   fmt.Println(dataset.reviews[1].ReviewText)
   fmt.Println(dataset.reviews[1].Tokens)
   fmt.Println(dataset.reviews[1].WordCount)
}