package federalregister

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func fetch(pubLTE, pubGTE sql.NullString) []NoticeResult {
	url := "https://www.federalregister.gov/api/v1/articles.json"
	url += "?per_page=1000&order=newest&conditions[type][]=RULE"
	url += "&fields[]=document_number&fields[]=publication_date"
	url += "&fields[]=agency_names"
	if pubLTE.Valid {
		url += "&conditions[publication_date][lte]=" + pubLTE.String
	}
	if pubGTE.Valid {
		url += "&conditions[publication_date][gte]=" + pubGTE.String
	}

	resp, err := http.Get(url)

	if err != nil {
		log.Panic(err)
	}

	var res NoticeResults
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&res); err != nil {
		log.Panic(err)
	}
	return res.Results
}
