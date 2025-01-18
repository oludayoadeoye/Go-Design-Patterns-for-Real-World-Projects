package main

import (
   "bufio"
   "encoding/json"
   "errors"
   "fmt"
   "log"
   "os"
   "regexp"
   "strings"
)

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
   Tokens         []string
   WordCount      map[string]int
}

type Dataset struct {
   filepath string
   reviews  []Review
}

func (dataset *Dataset) empty() bool {
   if len(dataset.reviews) == 0 {
      return true
   }
   return false
}

func (dataset *Dataset) read_json_file() (bool, error) {
   // read the json file using the os package
   content, err := os.Open(dataset.filepath)
   // if we have an error, we log the error and exit the program
   if err != nil {
      return true, err
   }

   // defer closing the file until the read_json_file function finishes
   defer content.Close()

   // create a scanner variable that we will use to iterate
   // through the reviews
   scanner := bufio.NewScanner(content)

   // split the content of the file based on lines (each line is a review)
   scanner.Split(bufio.ScanLines)

   for scanner.Scan() {
      // we can iterate through and display each review
      //fmt.Println(scanner.Text()) // This is commented; otherwise, it
      // will print the entire file, which
      // will take a while. Uncomment if you
      // want to see the content of the file.
      var review Review
      err := json.Unmarshal([]byte(scanner.Text()), &review)
      if err != nil {
         return true, err
      }
      dataset.reviews = append(dataset.reviews, review)
   }
   return false, nil
}

func tokenize(text string) []string {
   // Set up the regexp to use a punctuation list
   re := regexp.MustCompile(`[[:punct:]]`)

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

func (dataset *Dataset) tokenize() (bool, error) {
   if dataset.empty() {
      return true, errors.New("Dataset is empty. Please read data from json first.")
   }
   for i := range dataset.reviews {
      dataset.reviews[i].Tokens = tokenize(dataset.reviews[i].ReviewText)
   }
   return false, nil
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

func (dataset *Dataset) count_words() (bool, error) {
   if dataset.empty() {
      return true, errors.New("Dataset is empty. Please read data from json first.")
   }
   for i := range dataset.reviews {
      if len(dataset.reviews[i].Tokens) == 0 {
         dataset.reviews[i].Tokens = tokenize(dataset.reviews[i].ReviewText)
      }
      dataset.reviews[i].WordCount = count_words(dataset.reviews[i].Tokens)
   }
   return false, nil
}

func main() {
   dataset := Dataset{filepath: "./Digital_Music_5.json"}
   _, err := dataset.read_json_file()
   if err != nil {
      log.Fatal(err.Error())
   }

   _, err = dataset.tokenize()
   if err != nil {
      log.Fatal(err.Error())
   }

   dataset.count_words()

   fmt.Println(dataset.reviews[0].ReviewText)
   fmt.Println("---")
   fmt.Println(dataset.reviews[0].Tokens)
   fmt.Println("---")
   fmt.Println(dataset.reviews[0].WordCount)
}
