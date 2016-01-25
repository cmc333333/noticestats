package federalregister

import (
	"database/sql"
	"github.com/cmc333333/noticestats/db"
	"log"
)

func Backfill() {
	conn := db.NewConnection()
	var minPublished sql.NullString
	row := conn.QueryRow("SELECT min(published) from notice")
	if err := row.Scan(&minPublished); err != nil {
		log.Fatal(err)
	}

	results := fetch(minPublished, sql.NullString{"", false})
	for _, result := range results {
		insertIfNew(conn, &result)
	}
}
