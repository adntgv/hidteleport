package client

import (
	"fmt"
	"log"
	"sync"
)

type Client struct {
	Logger    *log.Logger
	WSClient  *WebSocketClient
	UDPClient *UDPClient
}

type Config struct {
	Logger                                                *log.Logger
	ServerHost, WSServerPath, WSServerPort, UDPServerPort string
	KeyboardInChan, MouseInChan                           chan []byte
}

func NewClient(c *Config) *Client {
	wsServerAddress := fmt.Sprintf("%v:%v", c.ServerHost, c.WSServerPort)
	udpServerAddress := fmt.Sprintf("%v:%v", c.ServerHost, c.UDPServerPort)
	return &Client{
		Logger:    c.Logger,
		WSClient:  NewWebSocketClient(c.Logger, wsServerAddress, c.WSServerPath, c.KeyboardInChan),
		UDPClient: &UDPClient{UDPServerAddress: udpServerAddress, InChan: c.MouseInChan},
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
