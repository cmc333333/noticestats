package db

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"github.com/jmoiron/sqlx"
)

func config(env string) (*goose.DBConf, error) {
	return goose.NewDBConf("db", env, "")
}

func NewConnection(env string) (*sqlx.DB, error) {
	dbconfig, err := config(env)
	if err != nil {
		return nil, err
	}
	return sqlx.Connect(dbconfig.Driver.Name, dbconfig.Driver.OpenStr)
}
