package main

import (
	"flag"
	"log"
	"os"
)

var (
	addr = flag.String("addr", "localhost:8080", "address to serve on / connect to")
)

func main() {
	flag.Parse()
	log.SetFlags(0)

	var action string
	for _, arg := range os.Args {
		if arg == "connect" {
			action = "connect"
		}
	}
	switch action {
	case "connect":
		log.Println("connecting to ", *addr)
		clientRun(*addr)
	default:
		log.Println("serving on ", *addr)
		serverRun(*addr)
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
