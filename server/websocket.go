package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

type WebSocketServer struct {
	Logger  *log.Logger
	Address string
	Path    string
	OutChan chan []byte
}

func NewWebSocketServer(logger *log.Logger, address, path string, outChan chan []byte) *WebSocketServer {
	return &WebSocketServer{
		Logger:  logger,
		Address: address,
		Path:    path,
		OutChan: outChan,
	}
}

func (s *WebSocketServer) Run() error {
	http.HandleFunc(s.Path, s.handler)
	return http.ListenAndServe(s.Address, nil)
}

func (s *WebSocketServer) handler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.Logger.Println("upgrade:", err)
		return
	}
	defer c.Close()
	for data := range s.OutChan {
		if err := c.WriteMessage(websocket.BinaryMessage, data); err != nil {
			s.Logger.Println("write:", err)
			break
		}
	}
}
