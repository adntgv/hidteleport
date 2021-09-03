package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/go-vgo/robotgo"
	"github.com/gorilla/websocket"
	hook "github.com/robotn/gohook"
)

var upgrader = websocket.Upgrader{} // use default options

type position struct {
	x int
	y int
}

func serverRun(address, udpAddress string, size screenSize) {
	var inputEvents = make(chan eventWrapper)

	go func(inputEvents chan eventWrapper) {
		EvChan := hook.Start()
		defer hook.End()

		p := position{
			x: int(size.width / 2),
			y: int(size.height / 2),
		}
		for ev := range EvChan {
			evtWpr, skip := toEventWrapper(ev, size, p)
			if skip {
				continue
			}
			inputEvents <- evtWpr
		}
	}(inputEvents)

	go func() {
		srv := &udpserver{
			inputDataChan: newDataChan(inputEvents),
			log:           *log.Default(),
			lock:          &sync.Mutex{},
		}
		srv.serveUDP(udpAddress)
	}()

	http.HandleFunc("/", getHandler(size, inputEvents))
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(address, nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("echo")); err != nil {
		log.Println(err)
	}
}

func getHandler(size screenSize, inputEvents chan eventWrapper) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()
		for wpr := range inputEvents {
			err := c.WriteJSON(wpr)
			if err != nil {
				log.Println("write:", err)
				break
			}

			robotgo.Move(int(size.width/2), int(size.height/2))
		}
	}
}

func toEventWrapper(evt hook.Event, size screenSize, p position) (eventWrapper, bool) {
	dta := delta{
		X: int(evt.X) - p.x,
		Y: int(evt.Y) - p.y,
	}
	if dta.X == 0 && dta.Y == 0 {
		return eventWrapper{}, true
	}
	smp := scaledPosition{
		X: float32(evt.X) / size.width,
		Y: float32(evt.Y) / size.height,
	}

	return eventWrapper{
		Event:               evt,
		ScaledMousePosition: smp,
		Delta:               dta,
	}, false
}

func newDataChan(anyChan chan eventWrapper) chan []byte {
	bzChan := make(chan []byte)
	go func(anyChan chan eventWrapper) {
		for any := range anyChan {
			bz, err := json.Marshal(any)
			if err != nil {
				log.Println(err)
				continue
			}
			bzChan <- bz
		}
	}(anyChan)
	return bzChan
}
