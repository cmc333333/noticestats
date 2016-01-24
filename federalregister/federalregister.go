package federalregister

import (
	"encoding/json"
	"github.com/cmc333333/noticestats/db"
	"log"
	"net/http"
	"time"
)

type NoticeResult struct {
	DocumentNumber string `json:"document_number"`
	Published      string `json:"publication_date"`
}

type NoticeResults struct {
	Count   int
	Results []NoticeResult
}

func Fetch(pubLTE time.Time) []NoticeResult {
	url := "https://www.federalregister.gov/api/v1/articles.json"
	url += "?per_page=1000&order=newest&conditions[type][]=RULE"
	url += "&conditions[publication_date][lte]="
	url += pubLTE.Format("2006-01-02")

	resp, err := http.Get(url)

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

func Sync() {
	conn := db.NewConnection()
	for _, result := range Fetch(time.Now()) {
		var count int
		row := conn.QueryRow(
			"SELECT COUNT(*) from notice WHERE id=?",
			result.DocumentNumber)
		if err := row.Scan(&count); err != nil {
			log.Fatal(err)
		}
		if count == 0 {
			sql := "INSERT INTO notice (id, published) "
			sql += "VALUES (?, ?)"
			conn.MustExec(
				sql, result.DocumentNumber, result.Published)
		}
	}
}
