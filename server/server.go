package server

import (
	"fmt"
	"sync"

	"github.com/adntgv/hidteleport/types"
	"go.uber.org/zap"
)

type Server struct {
	logger      *zap.Logger
	WSServer    *WebSocketServer
	Broadcaster *Broadcaster
}

type Config struct {
	types.Config
}

func NewServer(c *Config) *Server {
	wsServerAddress := fmt.Sprintf("%v:%v", c.Host, c.WSServerPort)
	broadcasterAddress := fmt.Sprintf("%v:%v", c.Host, c.BroadcasterPort)
	return &Server{
		logger:      c.Logger,
		WSServer:    NewWebSocketServer(c.Logger, wsServerAddress, c.WSServerPath, c.KeyboardChan),
		Broadcaster: NewBroadcaster(c.Logger, broadcasterAddress, c.MouseChan),
	}
}

func (s *Server) Run() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := s.WSServer.Run(); err != nil {
			s.logger.Sugar().Error(err)
		}
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := s.Broadcaster.Run(); err != nil {
			s.logger.Sugar().Error(err)
		}
	}(wg)

	wg.Wait()
}
