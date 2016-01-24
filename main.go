package main

import (
	"flag"
	"fmt"
	"github.com/cmc333333/noticestats/db"
	"github.com/cmc333333/noticestats/federalregister"
	"github.com/cmc333333/noticestats/web"
)

func main() {
	var command string
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		command = "web"
	} else {
		command = args[0]
	}

	switch command {
	case "web":
		web.Run()
	case "db-update":
		db.CreateOrUpdate("development")
	case "fr-sync":
		federalregister.Sync("development")
	default:
		fmt.Printf("Unknown command: %s", command)
	}
}
