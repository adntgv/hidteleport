package server

import (
	"net/http"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{} // use default options

type WebSocketServer struct {
	Logger  *zap.Logger
	Address string
	Path    string
	OutChan chan []byte
}

func NewWebSocketServer(logger *zap.Logger, address, path string, outChan chan []byte) *WebSocketServer {
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
		s.Logger.Sugar().Error("upgrade:", err)
		return
	}
	defer c.Close()
	for data := range s.OutChan {
		if err := c.WriteMessage(websocket.BinaryMessage, data); err != nil {
			s.Logger.Sugar().Error("write:", err)
			break
		}
	}
}
