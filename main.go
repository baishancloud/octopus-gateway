package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/baishancloud/octopux-gateway/g"
	"github.com/baishancloud/octopux-gateway/http"
	"github.com/baishancloud/octopux-gateway/receiver"
	"github.com/baishancloud/octopux-gateway/sender"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	// global config
	g.ParseConfig(*cfg)

	sender.Start()
	receiver.Start()

	// http
	http.Start()

	select {}
}
