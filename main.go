package main

import (
	"flag"
	"fmt"
	"github.com/cmc333333/noticestats/db"
	"github.com/cmc333333/noticestats/federalregister"
	"github.com/cmc333333/noticestats/stats"
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
		db.CreateOrUpdate()
	case "fr-backfill":
		federalregister.Backfill()
	case "fr-sync-new":
		federalregister.SyncNew()
	case "process-one":
		stats.Process()
	default:
		fmt.Printf("Unknown command: %s", command)
	}
}
