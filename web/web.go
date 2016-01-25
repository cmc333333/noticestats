package web

import (
	"database/sql"
	"fmt"
	"github.com/cmc333333/noticestats/db"
	"log"
	"net/http"
	"os"
)

func Run() {
	http.HandleFunc("/", hello)
	host := os.Getenv("OPENSHIFT_GO_IP")
	port := os.Getenv("OPENSHIFT_GO_PORT")
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "8080"
	}
	bind := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("listening on %s...", bind)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	conn := db.NewConnection()
	rows, err := conn.Query(`
		SELECT notice_id, agency
		FROM notice LEFT JOIN notice_agency ON (notice.id = notice_id)
		ORDER BY published, notice_id, agency`)
	if err != nil {
		log.Panic(err)
	}
	lastNotice := ""
	for rows.Next() {
		var docNum string
		var agency sql.NullString
		if err := rows.Scan(&docNum, &agency); err != nil {
			log.Fatal(err)
		}
		if docNum != lastNotice {
			fmt.Fprintf(res, "%s\n", docNum)
			lastNotice = docNum
		}
		if agency.Valid {
			fmt.Fprintf(res, "\t%s\n", agency.String)
		}
	}
}
