package main

import (
	"flag"
	"fmt"
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
	default:
		fmt.Printf("Unknown command: %s", command)
	}
}
