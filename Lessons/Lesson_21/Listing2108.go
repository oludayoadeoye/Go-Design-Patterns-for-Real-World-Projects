func main() {
    reviews := read_json_file("./Digital_Music_5.json")
    fmt.Println(reviews[0].ReviewText)
    fmt.Println("-- - - - - - - - - ")
    fmt.Println(reviews[1].ReviewText)
}