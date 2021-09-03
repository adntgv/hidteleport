package server

import (
	"fmt"
	"log"
	"sync"
)

type Server struct {
	logger      *log.Logger
	WSServer    *WebSocketServer
	Broadcaster *Broadcaster
}

type Config struct {
	Logger                                            *log.Logger
	Host, WSServerPort, WSServerPath, BroadcasterPort string
	WSServerOutChan, BroadcasterOutChan               chan []byte
}

func NewServer(c *Config) *Server {
	wsServerAddress := fmt.Sprintf("%v:%v", c.Host, c.WSServerPort)
	broadcasterAddress := fmt.Sprintf("%v:%v", c.Host, c.BroadcasterPort)
	return &Server{
		logger:      c.Logger,
		WSServer:    NewWebSocketServer(c.Logger, wsServerAddress, c.WSServerPath, c.WSServerOutChan),
		Broadcaster: NewBroadcaster(c.Logger, broadcasterAddress, c.BroadcasterOutChan),
	}
}

func (s *Server) Run() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := s.WSServer.Run(); err != nil {
			s.logger.Println(err)
		}
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := s.WSServer.Run(); err != nil {
			s.logger.Println(err)
		}
	}(wg)

	wg.Wait()
}
