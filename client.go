package main

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func clientRun(address string) {
	u := url.URL{Scheme: "ws", Host: address, Path: "/"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	log.Printf("conneted to %s", u.String())

	for {
		_, bz, err := c.ReadMessage()
		if err != nil {
			log.Println("ReadMessage:", err)
			return
		}
		message, err := eventsFromBytesJSON(bz)
		if err != nil {
			log.Println("eventsFromBytes:", err)
			return
		}

		log.Printf("recv: %s", message)
	}
}
