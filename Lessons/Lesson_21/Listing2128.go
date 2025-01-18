func main() {
   dataset := Dataset{filepath: "./Digital_Music_5.json"}
   _, err := dataset.tokenize()
   if err != nil {
      log.Fatal(err.Error())
   }
}