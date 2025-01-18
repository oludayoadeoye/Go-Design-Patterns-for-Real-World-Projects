func main() {
   reviews := read_json_file("./Digital_Music_5.json")
   tokens := tokenize(reviews[0].ReviewText)
   fmt.Print("tokens: ")
   fmt.Println(tokens)
}