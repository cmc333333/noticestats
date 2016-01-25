package federalregister

import "github.com/jmoiron/sqlx"

func insertNotices(conn *sqlx.DB, results []NoticeResult) {
	sql := conn.Rebind(`
		INSERT INTO notice(id, published)
		SELECT ?, ? 
		WHERE NOT EXISTS (SELECT 1 FROM notice WHERE id = ?)`)
	for _, result := range results {
		conn.MustExec(sql, result.DocumentNumber, result.Published, result.DocumentNumber)
	}
	insertAgencies(conn, results)
}

func insertAgencies(conn *sqlx.DB, results []NoticeResult) {
	delete := conn.Rebind("DELETE FROM notice_agency WHERE notice_id = ?")
	insert := conn.Rebind(`
		INSERT INTO notice_agency (notice_id, agency)
		VALUES (?, ?)`)
	for _, result := range results {
		conn.MustExec(delete, result.DocumentNumber)
		for _, agency := range result.Agencies {
			conn.MustExec(insert, result.DocumentNumber, agency)
		}
	}
}
