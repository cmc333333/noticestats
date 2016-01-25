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
		log.Panic(err)
	}

	insertNotices(conn, fetch(sql.NullString{"", false}, maxPublished))
}
