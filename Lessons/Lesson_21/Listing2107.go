func read_json_file(filepath string) []Review {
    // read the json file using the os package
    content, err := os.Open(filepath)
    // if we have an error, we log the error and exit the program
    if err != nil {
        log.Fatal(err)
    }
    // defer closing the file until the read_json_file function finishes
    defer content.Close()
    // create a scanner variable that we will use to iterate through the reviews
    scanner := bufio.NewScanner(content)
    // split the content of the file based on lines (each line is a review)
    scanner.Split(bufio.ScanLines)
    var reviews []Review
    for scanner.Scan() {
        // We can iterate through each review
        var review Review
        err := json.Unmarshal([]byte(scanner.Text()), &review)
        if err != nil {
            log.Fatal(err)
        }
        reviews = append(reviews, review)
    }
    return reviews
}