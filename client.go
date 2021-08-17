package main

import (
	"log"
	"net/url"

	"github.com/go-vgo/robotgo"
	"github.com/gorilla/websocket"
	hook "github.com/robotn/gohook"
)

func clientRun(address string, size screenSize) {
	u := url.URL{Scheme: "ws", Host: address, Path: "/"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	log.Printf("conneted to %s", u.String())

	for {
		evtWpr := eventWrapper{
			Event:               hook.Event{},
			ScaledMousePosition: scaledPosition{},
		}
		err := c.ReadJSON(&evtWpr)
		if err != nil {
			log.Println("ReadMessage:", err)
			continue
		}
		if err := apply(evtWpr, size); err != nil {
			log.Println(err)
			continue
		}
	}
}

func apply(evtWpr eventWrapper, size screenSize) error {
	//log.Println(evtWpr)
	switch evtWpr.Kind {
	// Button, Clicks, X, Y, Amount, Rotation and Direction
	case hook.MouseDown,
		hook.MouseUp,
		hook.MouseHold,
		hook.MouseDrag,
		hook.MouseWheel:
		// stub
	case hook.MouseMove:
		robotgo.MoveRelative(evtWpr.Delta.X, evtWpr.Delta.Y)

		// Mask, Keycode, Rawcode, and Keychar,
		// Keychar is probably what you want.
	case hook.KeyDown,
		hook.KeyHold,
		hook.KeyUp:

	}

	return nil
}
