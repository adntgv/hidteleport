package client

import (
	"fmt"
	"sync"

	"github.com/adntgv/hidteleport/types"
	"go.uber.org/zap"
)

type Client struct {
	Logger    *zap.Logger
	WSClient  *WebSocketClient
	UDPClient *UDPClient
}

type Config struct {
	types.Config
}

func NewClient(c *Config) *Client {
	wsServerAddress := fmt.Sprintf("%v:%v", c.Host, c.WSServerPort)
	udpServerAddress := fmt.Sprintf("%v:%v", c.Host, c.BroadcasterPort)
	return &Client{
		Logger:    c.Logger.Named("client"),
		WSClient:  NewWebSocketClient(c.Logger.Named("wsclient"), wsServerAddress, c.WSServerPath, c.KeyboardChan),
		UDPClient: &UDPClient{logger: c.Logger.Named("udpclient"), UDPServerAddress: udpServerAddress, InChan: c.MouseChan},
	}
}

func (client *Client) Run() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := client.WSClient.Run(); err != nil {
			client.Logger.Sugar().Error(err)
		}
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := client.UDPClient.Run(); err != nil {
			client.Logger.Sugar().Error(err)
		}
	}(wg)

	wg.Wait()
}
