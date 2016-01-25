package stats

import (
	"database/sql"
	"encoding/json"
	"github.com/cmc333333/noticestats/db"
	"log"
	"net/http"
)

func Process() {
	conn := db.NewConnection()
	var nextDoc sql.NullString
	row := conn.QueryRow(`
		SELECT notice.id
		FROM notice LEFT JOIN stats ON (notice.id = stats.id)
		WHERE stats.id IS NULL
		LIMIT 1`)
	if err := row.Scan(&nextDoc); err != nil {
		log.Panic(err)
	}

	if nextDoc.Valid {
		notice := fetch(nextDoc.String)
		xml := fetchXML(&notice)
		sql := conn.Rebind(`
			INSERT INTO stats
			(id, xml_len, regtext_len, page_len, is_correction)
			VALUES (?, ?, ?, ?, ?)`)
		conn.MustExec(
			sql,
			nextDoc.String, len(xml), regtextLen(xml),
			notice.PageLength, notice.CorrectionOf != "")

	}
}

func fetch(docNum string) NoticeJSON {
	url := "https://www.federalregister.gov/api/v1/articles/" + docNum
	url += ".json?fields[]=full_text_xml_url&fields[]=page_length"
	url += "&fields[]=correction_of"

	resp, err := http.Get(url)

	if err != nil {
		log.Panic(err)
	}

	var result NoticeJSON
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		log.Panic(err)
	}
	return result
}
