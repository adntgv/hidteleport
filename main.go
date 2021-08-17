package main

import (
	"flag"
	"log"
)

func main() {
	flag.Parse()
	log.SetFlags(0)

	switch detectRunMode(flag.Args()) {
	case "client":
		clientRun(*connect)
	default:
		serverRun(*serve)
	}
}

func detectRunMode(args []string) string {
	for _, arg := range args {
		if arg == "serve" {
			return "server"
		} else if arg == "connect" {
			return "client"
		}
	}

	return ""
}
