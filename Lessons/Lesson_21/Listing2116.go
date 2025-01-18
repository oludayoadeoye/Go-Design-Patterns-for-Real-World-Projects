package main

import (
   "bufio"
   "encoding/json"
   "log"
   "os"
   "regexp"
   "strings"
)

func main() {
   reviews := read_json_file("./Digital_Music_5.json")
   for i := range reviews {
      tokens := tokenize(reviews[i].ReviewText) // tokenize review
      count_words(tokens)
   }
}

type Review struct {
   ReviewerID     string  `json:"reviewerID"`
   Asin           string  `json:"asin"`
   ReviewerName   string  `json:"reviewerName"`
   Helpful        [2]int  `json:"helpful"`
   ReviewText     string  `json:"reviewText"`
   Overall        float32 `json:"overall"`
   Summary        string  `json:"summary"`
   UnixReviewTime int64   `json:"unixReviewTime"`
   ReviewTime     string  `json:"reviewTime"`
}

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

func tokenize(text string) []string {

   // Set up the regexp to use a punctuation list
   re := regexp.MustCompile("[.,!?\\-_#^()+=;/&'\"]")

   // use the ReplaceAllString to replace any punctuation with a space
   w := re.ReplaceAllString(text, " ")

   w = strings.ToLower(w) // convert to lowercase

   // Use the Fields function from the strings package to split
   // the string w around each instance of one or more consecutive
   // whitespace characters
   tokens := strings.Fields(w)

   // return the slice, which represents the list of tokens in
   // order from the input string
   return tokens
}

func count_words(words []string) map[string]int {

   word_count := make(map[string]int)

   for i := range words {
      if _, ok := word_count[words[i]]; ok {
         word_count[words[i]] = word_count[words[i]] + 1
      } else {
         word_count[words[i]] = 1
      }
   }

   return word_count
}
