package db

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"log"
)

func CreateOrUpdate() {
	conf, err := config()
	if err != nil {
		log.Panic(err)
	}
	dir := conf.MigrationsDir
	target, err := goose.GetMostRecentDBVersion(dir)
	if err != nil {
		log.Panic(err)
	}
	if err := goose.RunMigrations(conf, dir, target); err != nil {
		log.Panic(err)
	}
}
