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

func clientRun(address, udpaddr string, useUDP bool, size screenSize) {
	u := url.URL{Scheme: "ws", Host: address, Path: "/"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	log.Printf("conneted to %s", u.String())

	if useUDP {
		go handleUDP(udpaddr, size)
	}

	for {
		evtWpr, err := receiveEventWrapper(c)
		if err != nil {
			log.Println(err)
			continue
		}

		if err := applyKey(evtWpr, size); err != nil {
			log.Println(err)
			continue
		}
	}
}

func applyKey(evtWpr eventWrapper, size screenSize) error {
	//log.Println(evtWpr)
	switch evtWpr.Kind {
	case hook.KeyUp, hook.KeyDown:
		log.Println("kdown")
		robotgo.KeyTap(string(evtWpr.Keychar))
	}

	return nil
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

var mouseIsDown = false

func getKey(keycode rune) string {
	akey, ok := keysMap[uint16(keycode)]
	if !ok {
		log.Printf("%v not found in map\n", keycode)
	}
	log.Println(keycode)
	log.Println(akey)
	return akey
}

func handleUDP(addr string, size screenSize) {
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

		evtWpr := eventWrapper{
			Event:               hook.Event{},
			ScaledMousePosition: scaledPosition{},
		}
		err = json.Unmarshal(buffer[:n], &evtWpr)
		if err != nil {
			log.Println(err)
			continue
		}

		if err := applyMouse(evtWpr, size); err != nil {
			log.Println(err)
			continue
		}
	}
}

func applyMouse(evtWpr eventWrapper, size screenSize) error {
	//log.Println(evtWpr)
	switch evtWpr.Kind {
	case hook.MouseMove:
		log.Println("move")
		robotgo.MoveRelative(evtWpr.Delta.X, evtWpr.Delta.Y)
	case hook.MouseDown:
		log.Println("mdown")
		mouseIsDown = true
	case hook.MouseUp:
		log.Println("mup")
		if mouseIsDown {
			robotgo.Click()
		}
		mouseIsDown = false
	case hook.MouseDrag:
		log.Println("mdrag")
		robotgo.DragMouse(evtWpr.Delta.X, evtWpr.Delta.Y)
	}
	return nil
}
