package db

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"github.com/jmoiron/sqlx"
	"log"
)

func config(env string) (*goose.DBConf, error) {
	return goose.NewDBConf("db", env, "")
}

func NewConnection(env string) *sqlx.DB {
	dbconfig, err := config(env)
	if err != nil {
		log.Fatal(err)
	}
	return sqlx.MustConnect(dbconfig.Driver.Name, dbconfig.Driver.OpenStr)
}
