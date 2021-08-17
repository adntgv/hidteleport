package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	hook "github.com/robotn/gohook"
)

var upgrader = websocket.Upgrader{} // use default options
var inputEvents = make(chan hook.Event)

func serverRun(address string) {
	go func(inputEvents chan hook.Event) {
		EvChan := hook.Start()
		defer hook.End()

		for ev := range EvChan {
			inputEvents <- ev
		}
	}(inputEvents)

	http.HandleFunc("/", handle)
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(address, nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("echo")); err != nil {
		log.Println(err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for evt := range inputEvents {
		/*bz, err := eventToBytesJSON(evt)
		if err != nil {
			log.Println("write:", err)
			continue
		}
		err = c.WriteMessage(websocket.BinaryMessage, bz)*/
		err := c.WriteJSON(evt)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
