package main

import (
	"flag"
	"fmt"

	"github.com/vpoliakov01/CoverageMonitor/back_end/server"
)

func main() {
	cfg := server.Config{}
	flag.IntVar(&cfg.Port, "port", 8080, "serving port")
	flag.BoolVar(&cfg.Mock, "mock", false, "replies with a mock response for get info to avoid github api rate limits")
	flag.Parse()

	if cfg.Mock {
		fmt.Println("RUNNING IN MOCK MODE")
	}

	server := server.New(&cfg)
	err := server.Serve()
	if err != nil {
		panic(err)
	}
}
