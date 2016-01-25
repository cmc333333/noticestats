package federalregister

import (
	"database/sql"
	"github.com/cmc333333/noticestats/db"
	"log"
)

func SyncNew() {
	conn := db.NewConnection()
	var maxPublished sql.NullString
	row := conn.QueryRow("SELECT max(published) from notice")
	if err := row.Scan(&maxPublished); err != nil {
		log.Fatal(err)
	}

	results := fetch(sql.NullString{"", false}, maxPublished)
	for _, result := range results {
		insertIfNew(conn, &result)
	}
}
