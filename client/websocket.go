package client

import (
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type WebSocketClient struct {
	logger          *zap.Logger
	WSServerAddress string
	WSServerPath    string
	InChan          chan []byte
}

func NewWebSocketClient(logger *zap.Logger, wsServerAddress, wsServerPath string, inChan chan []byte) *WebSocketClient {
	return &WebSocketClient{
		logger:          logger.Named("wsclient"),
		WSServerAddress: wsServerAddress,
		WSServerPath:    wsServerPath,
		InChan:          inChan,
	}
}

func (client *WebSocketClient) Run() error {
	u := url.URL{Scheme: "ws", Host: client.WSServerAddress, Path: client.WSServerPath}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return fmt.Errorf("dial: %v", err)
	}
	defer c.Close()

	client.logger.Sugar().Infof("conneted to %s", u.String())

	for {
		_, bz, err := c.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err) {
				return nil
			}
			client.logger.Sugar().Errorf("read: %v", err)
			continue
		}

		client.InChan <- bz
	}
}
