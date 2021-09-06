package main

import (
	"flag"
	"sync"

	"log"

	"github.com/adntgv/hidteleport/client"
	"github.com/adntgv/hidteleport/emulator"
	"github.com/adntgv/hidteleport/events"
	"github.com/adntgv/hidteleport/server"
	"github.com/adntgv/hidteleport/types"
)

var (
	mode          = flag.String("mode", "server", "in which mode to run")
	host          = flag.String("host", "localhost", "server host address")
	wsServerPath  = flag.String("path", "/ws", "websocket server path")
	wsServerPort  = flag.String("ws-port", "8080", "websocket serverport for")
	udpServerPort = flag.String("udp-port", "8081", "udp server port")
)

func main() {
	flag.Parse()

	logger := log.Default()
	screen := &types.Screen{} // Needed for absolute mouse positioning
	wg := &sync.WaitGroup{}
	keyboardChan := make(chan []byte)
	mouseChan := make(chan []byte)

	commonConfig := types.Config{
		Logger:          logger,
		Host:            *host,
		WSServerPath:    *wsServerPath,
		WSServerPort:    *wsServerPort,
		BroadcasterPort: *udpServerPort,
		KeyboardChan:    keyboardChan,
		MouseChan:       mouseChan}

	if *mode == "client" {
		wg.Add(2)
		client := client.NewClient(&client.Config{Config: commonConfig})
		go client.Run()

		emulator := emulator.NewEmulator(logger, screen, mouseChan)
		go emulator.Run()
	} else {
		wg.Add(2)
		transformer := events.NewTransformer(&types.Coordinates{}, screen)
		producer := events.NewProducer(transformer, logger, mouseChan, keyboardChan)
		go producer.Run()

		server := server.NewServer(&server.Config{Config: commonConfig})
		go server.Run()
	}

	wg.Wait()
}
