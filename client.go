package main

import (
	"encoding/json"
	"log"
	"net"
	"net/url"

	"github.com/go-vgo/robotgo"
	"github.com/gorilla/websocket"
	hook "github.com/robotn/gohook"
)

func clientRun(address, udpaddr string, size screenSize) {
	u := url.URL{Scheme: "ws", Host: address, Path: "/"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	log.Printf("conneted to %s", u.String())

	go handleUDP(udpaddr)

	for {
		_, err := receiveEventWrapper(c)
		if err != nil {
			log.Println(err)
			continue
		}

		/*if err := apply(evtWpr, size); err != nil {
			log.Println(err)
			continue
		}*/
	}
}

func receiveEventWrapper(c *websocket.Conn) (eventWrapper, error) {
	evtWpr := eventWrapper{
		Event:               hook.Event{},
		ScaledMousePosition: scaledPosition{},
	}
	err := c.ReadJSON(&evtWpr)
	if err != nil {
		return evtWpr, err
	}
	return evtWpr, nil
}

func apply(evtWpr eventWrapper, size screenSize) error {
	//log.Println(evtWpr)
	switch evtWpr.Kind {
	case hook.MouseMove:
		robotgo.MoveRelative(evtWpr.Delta.X, evtWpr.Delta.Y)
	}

	return nil
}

func handleUDP(addr string) {
	s, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Println(err)
		return
	}

	connection, err := net.DialUDP("udp", nil, s)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("udp client connected to", addr)

	defer connection.Close()

	_, err = connection.Write([]byte("hi from " + connection.LocalAddr().String()))
	if err != nil {
		log.Println(err)
		return
	}

	buffer := make([]byte, 1024)
	for {
		n, _, err := connection.ReadFromUDP(buffer)
		if err != nil {
			log.Println(err)
			continue
		}

		evt := hook.Event{}
		err = json.Unmarshal(buffer[:n], &evt)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println(evt)
	}
}
