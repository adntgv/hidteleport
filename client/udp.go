package client

import (
	"fmt"
	"log"
	"net"
)

type UDPClient struct {
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

	log.Println("udp client connected to", client.UDPServerAddress)

	defer connection.Close()

	_, err = connection.WriteToUDP([]byte("hi from "+connection.LocalAddr().String()), s)
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
