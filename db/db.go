package db

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"flag"
	"github.com/jmoiron/sqlx"
	"log"
)

var env = flag.String("env", "development", "which DB environment to use")

func config() (*goose.DBConf, error) {
	return goose.NewDBConf("db", *env, "")
}

func NewConnection() *sqlx.DB {
	dbconfig, err := config()
	if err != nil {
		log.Panic(err)
	}
	return sqlx.MustConnect(dbconfig.Driver.Name, dbconfig.Driver.OpenStr)
}
