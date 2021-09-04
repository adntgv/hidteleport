package client

import (
	"fmt"
	"log"
	"sync"

	"github.com/adntgv/hidteleport/types"
)

type Client struct {
	Logger    *log.Logger
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
		Logger:    c.Logger,
		WSClient:  NewWebSocketClient(c.Logger, wsServerAddress, c.WSServerPath, c.KeyboardChan),
		UDPClient: &UDPClient{UDPServerAddress: udpServerAddress, InChan: c.MouseChan},
	}
}

func (client *Client) Run() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := client.WSClient.Run(); err != nil {
			client.Logger.Println(err)
		}
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := client.UDPClient.Run(); err != nil {
			client.Logger.Println(err)
		}
	}(wg)

	wg.Wait()
}
