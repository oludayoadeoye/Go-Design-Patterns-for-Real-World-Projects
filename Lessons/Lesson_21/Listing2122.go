func main() {
   dataset := Dataset{filepath: "./Digital_Music_5.json"}
   dataset.read_json_file()
   dataset.tokenize()
   fmt.Println(dataset.reviews[1].ReviewText)
   fmt.Println("---")
   fmt.Println(dataset.reviews[1].Tokens)
   fmt.Println("---")
   fmt.Println(dataset.reviews[2].ReviewText)
   fmt.Println("---")
   fmt.Println(dataset.reviews[2].Tokens)
}