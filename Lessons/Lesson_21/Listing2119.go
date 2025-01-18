func (dataset *Dataset) read_json_file() {
   // read the json file using the os package
   content, err := os.Open(dataset.filepath)
   // if we have an error, we log the error and exit the program
   if err != nil {
      log.Fatal(err)
   }

 // defer closing the file until the read_json_file function finishes.
   defer content.Close()

 // create a scanner variable that we will use to iterate 
 // through the reviews
   scanner := bufio.NewScanner(content)
 
// split the content of the file based on lines (each line is a review)
   scanner.Split(bufio.ScanLines)
   for scanner.Scan() {
      // we can iterate through and display each review
      // fmt.Println(scanner.Text()) // This is commented otherwise, it 
      // will print the entire file, which will take a while. Uncomment 
      // if you want to see the content of the file
      var review Review
      err := json.Unmarshal([]byte(scanner.Text()), &review)
      if err != nil {
         log.Fatal(err)
      }
   dataset.reviews = append(dataset.reviews, review)
   }
}