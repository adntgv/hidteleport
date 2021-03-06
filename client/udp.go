package client

import (
	"fmt"
	"net"

	"go.uber.org/zap"
)

type UDPClient struct {
	logger           *zap.Logger
	UDPServerAddress string
	InChan           chan []byte
}

func (client *UDPClient) Run() error {
	s, err := net.ResolveUDPAddr("udp", client.UDPServerAddress)
	if err != nil {
		return fmt.Errorf("resolve: %v", err)
	}

	connection, err := net.DialUDP("udp", nil, s)
	if err != nil {
		return fmt.Errorf("dial udp: %v", err)
	}

	client.logger.Sugar().Infof("udp client connected to %v", client.UDPServerAddress)

	defer connection.Close()

	_, err = connection.Write([]byte("hi from " + connection.LocalAddr().String()))
	if err != nil {
		return fmt.Errorf("write: %v", err)
	}

	buffer := make([]byte, 1024)
	for {
		n, _, err := connection.ReadFromUDP(buffer)
		if err != nil {
			return fmt.Errorf("read: %v", err)
		}

		client.InChan <- buffer[:n]
	}
}
