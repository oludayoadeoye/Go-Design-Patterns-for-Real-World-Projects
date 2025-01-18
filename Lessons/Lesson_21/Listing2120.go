func main() {
   dataset := Dataset{filepath: "./Digital_Music_5.json"}
   dataset.read_json_file()
   fmt.Println(dataset.reviews[1].ReviewText)
   fmt.Println(dataset.reviews[2].ReviewText)
}