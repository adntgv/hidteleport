package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-vgo/robotgo"
)

var (
	ip      = flag.String("ip", "localhost", "address to serve on / connect to")
	port    = flag.String("port", "8888", "http port to serve on / connect to")
	udpport = flag.String("udpport", "8889", "udp port to serve on / connect to")
	useUDP  = flag.Bool("udp", true, "should client use UDP connection rather than tcp")
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
	addr := fmt.Sprintf("%v:%v", *ip, *port)
	udpaddr := fmt.Sprintf("%v:%v", *ip, *udpport)
	switch action {
	case "connect":
		log.Println("connecting to ", addr)
		clientRun(addr, udpaddr, *useUDP, size)
	default:
		log.Println("serving on ", addr)
		serverRun(addr, udpaddr, size)
	}
}
