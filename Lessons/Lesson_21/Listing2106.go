package main

import (
   "fmt"
   "os"
   "bufio"
   "encoding/json"
   "log"
)

func main() {
    read_json_file("./Digital_Music_5.json")
}

type Review struct {    ReviewerID     string  `json:"reviewerID"`    Asin           string  `json:"asin"`    ReviewerName   string  `json:"reviewerName"`    Helpful        [2]int  `json:"helpful"`    ReviewText     string  `json:"reviewText"`    Overall        float32 `json:"overall"`    Summary        string  `json:"summary"`    UnixReviewTime int64   `json:"unixReviewTime"`    ReviewTime     string  `json:"reviewTime"`}func read_json_file(filepath string) {    // read the json file using the os package    content, err := os.Open(filepath)    // if we have an error, we log the error and exit the program    if err != nil {        log.Fatal(err)    }    // defer closing the file until the read_json_file function finishes.    defer content.Close()    // create a scanner variable that we will use to iterate through the reviews    scanner := bufio.NewScanner(content)    // split the content of the file based on lines (each line is a review)    scanner.Split(bufio.ScanLines)    for scanner.Scan() {        // We can iterate through and display each review        //fmt.Println(scanner.Text()) // Remove comment to print review line         var review Review        err := json.Unmarshal([]byte(scanner.Text()), &review)        if err != nil {            log.Fatal(err)            return        }        fmt.Println(review.Asin)    }}