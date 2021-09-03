package main

import (
	"flag"

	"log"

	"github.com/adntgv/hidteleport/client"
	"github.com/adntgv/hidteleport/server"
)

var (
	mode = flag.String("mode", "server", "in which mode to run")
)

func main() {
	flag.Parse()

	logger := &log.Logger{}

	if *mode == "client" {
		client := client.NewClient(&client.Config{
			Logger: logger,
		})
		client.Run()
	} else {
		server := server.NewServer(&server.Config{
			Logger: logger,
		})
		server.Run()
	}
}
