package web

import (
	"fmt"
	"github.com/cmc333333/noticestats/db"
	"github.com/cmc333333/noticestats/models"
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
	rows, err := conn.Queryx("SELECT * FROM notice ORDER BY published, id")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var notice models.Notice
		if err := rows.StructScan(&notice); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(res, "%s\n", notice)
	}
}
