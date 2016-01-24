package federalregister

import (
	"encoding/json"
	"log"
	"net/http"
)

type NoticeResult struct {
	DocumentNumber string `json:"document_number"`
}

type NoticeResults struct {
	Count   int
	Results []NoticeResult
}

func Fetch() []NoticeResult {
	resp, err := http.Get(
		"https://www.federalregister.gov/api/v1/articles.json")

	if err != nil {
		log.Print(err)
	}

	var res NoticeResults
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&res); err != nil {
		log.Fatal(err)
	}
	return res.Results
}
