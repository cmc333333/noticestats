package db

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"log"
)

func CreateOrUpdate(env string) {
	conf, err := config(env)
	if err != nil {
		log.Fatal(err)
	}
	dir := conf.MigrationsDir
	target, err := goose.GetMostRecentDBVersion(dir)
	if err != nil {
		log.Fatal(err)
	}
	if err := goose.RunMigrations(conf, dir, target); err != nil {
		log.Fatal(err)
	}
}
