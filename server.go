package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	hook "github.com/robotn/gohook"
)

func serverRun(address string) {
	go func(inputEvents chan hook.Event) {
		EvChan := hook.Start()
		defer hook.End()

		for ev := range EvChan {
			fmt.Println("hook: ", ev)
			inputEvents <- ev
		}
	}(inputEvents)

	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(address, nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for evt := range inputEvents {
		bz, err := eventToBytesJSON(evt)
		if err != nil {
			log.Println("write:", err)
			continue
		}
		err = c.WriteMessage(websocket.BinaryMessage, bz)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
