package main

import (
	"log"
	"net/url"

	"github.com/go-vgo/robotgo"
	"github.com/gorilla/websocket"
	hook "github.com/robotn/gohook"
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
			continue
		}
		event, err := eventsFromBytesJSON(bz)
		if err != nil {
			log.Println("eventsFromBytes:", err)
			continue
		}

		if err := apply(event); err != nil {
			log.Println(err)
			continue
		}
	}
}

func apply(evt hook.Event) error {
	switch evt.Kind {
	// Button, Clicks, X, Y, Amount, Rotation and Direction
	case hook.MouseDown,
		hook.MouseUp,
		hook.MouseHold,
		hook.MouseDrag,
		hook.MouseWheel:
		// stub
	case hook.MouseMove:
		robotgo.MoveMouse(int(evt.X), int(evt.Y))

		// Mask, Keycode, Rawcode, and Keychar,
		// Keychar is probably what you want.
	case hook.KeyDown,
		hook.KeyHold,
		hook.KeyUp:

	}

	return nil
}
