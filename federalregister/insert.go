package federalregister

import (
	"github.com/jmoiron/sqlx"
	"log"
)

func insert(conn *sqlx.DB, model *NoticeResult) {
	sql := "INSERT INTO notice(id, published) VALUES (?, ?)"
	conn.MustExec(sql, model.DocumentNumber, model.Published)

}

func insertIfNew(conn *sqlx.DB, model *NoticeResult) {
	var count int
	row := conn.QueryRow(
		"SELECT COUNT(*) FROM notice WHERE id=?",
		model.DocumentNumber)
	if err := row.Scan(&count); err != nil {
		log.Fatal(err)
	}
	if count == 0 {
		insert(conn, model)
	}
}
