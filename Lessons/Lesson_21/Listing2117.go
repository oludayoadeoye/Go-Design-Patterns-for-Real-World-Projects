type Review struct {
   ReviewerID string `json:"reviewerID"`
   Asin string `json:"asin"`
   ReviewerName string `json:"reviewerName"`
   Helpful [2]int `json:"helpful"`
   ReviewText string `json:"reviewText"`
   Overall float32 `json:"overall"`
   Summary string `json:"summary"`
   UnixReviewTime int64 `json:"unixReviewTime"`
   ReviewTime string `json:"reviewTime"`
   Tokens []string
   WordCount map[string]int
}