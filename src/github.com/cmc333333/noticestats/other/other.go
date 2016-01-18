package main

import (
	_ "encoding/json"
	"fmt"
	"github.com/cmc333333/noticestats/db"
	"github.com/cmc333333/noticestats/models"
	_ "net/http"
)

type NoticeResult struct {
	Url            string   `json:"full_text_xml_url"`
	DocumentNumber string   `json:"document_number"`
	Agencies       []string `json:"agency_names"`
}

type NoticeResults struct {
	Count       int
	NextPageUrl string `json:"next_page_url"`
	Results     []NoticeResult
}

func main() {
	db, err := db.NewConnection("development")
	if err != nil {
		panic(err)
	}

	vv := []models.Version{}
	err = db.Select(&vv, "SELECT id, version_id from goose_db_version")
	if err != nil {
		panic(err)
	}
	for _, v := range vv {
		fmt.Println(v)
	}
}
