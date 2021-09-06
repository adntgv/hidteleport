package server

import (
	"net"
	"sync"

	"go.uber.org/zap"
)

type Broadcaster struct {
	logger      *zap.Logger
	Address     string
	OutChan     chan []byte
	udpListener *net.UDPConn
	clientAddrs []*net.UDPAddr
	lock        *sync.Mutex
}

func NewBroadcaster(logger *zap.Logger, address string, outChan chan []byte) *Broadcaster {
	return &Broadcaster{
		logger:      logger.Named("broadcaster"),
		Address:     address,
		OutChan:     outChan,
		clientAddrs: make([]*net.UDPAddr, 0),
		lock:        &sync.Mutex{},
	}
}

func (b *Broadcaster) Run() error {
	var err error
	b.udpListener, err = newUDPListener(b.Address)
	if err != nil {
		return err
	}

	defer b.udpListener.Close()

	b.logger.Sugar().Infof("serving udp %v", b.Address)

	go b.broadcast()

	buffer := make([]byte, 1024)
	for {
		_, addr, err := b.udpListener.ReadFromUDP(buffer)
		if err != nil {
			b.logger.Sugar().Error(err)
			continue
		}

		b.logger.Sugar().Info("new udp client connected", addr.String(), string(buffer))
		b.addClientAddr(addr)
	}
}

func newUDPListener(address string) (*net.UDPConn, error) {
	s, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, err
	}

	return net.ListenUDP("udp", s)
}

func (b *Broadcaster) broadcast() {
	for data := range b.OutChan {
		for _, addr := range b.clientAddrs {
			if _, err := b.udpListener.WriteToUDP(data, addr); err != nil {
				b.logger.Sugar().Error(err)
				b.removeClientAddr(addr)
			}
		}
	}
}

func (b *Broadcaster) addClientAddr(addr *net.UDPAddr) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.clientAddrs = append(b.clientAddrs, addr)
}

func (b *Broadcaster) removeClientAddr(addr *net.UDPAddr) {
	b.lock.Lock()
	defer b.lock.Unlock()
	for i, clientAddr := range b.clientAddrs {
		if addr == clientAddr {
			b.clientAddrs = append(b.clientAddrs[:i], b.clientAddrs[i+1:]...)
		}
	}
}
