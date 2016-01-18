package db

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"github.com/jmoiron/sqlx"
)

func NewConnection(env string) (*sqlx.DB, error) {
	dbconfig, err := goose.NewDBConf("db", env, "")
	if err != nil {
		return nil, err
	}
	return sqlx.Connect(dbconfig.Driver.Name, dbconfig.Driver.OpenStr)
}
