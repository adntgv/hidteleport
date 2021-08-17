package main

import (
	"flag"
	"log"
	"os"

	"github.com/go-vgo/robotgo"
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

	x, y := robotgo.GetScaleSize()

	size := screenSize{
		width:  float32(x),
		height: float32(y),
	}

	switch action {
	case "connect":
		log.Println("connecting to ", *addr)
		clientRun(*addr, size)
	default:
		log.Println("serving on ", *addr)
		serverRun(*addr, size)
	}
}
