package main

import (
	"log"
	"net"
	"sync"
)

type udpserver struct {
	inputDataChan chan []byte
	udpListener   *net.UDPConn
	log           log.Logger
	clientAddrs   []*net.UDPAddr
	lock          *sync.Mutex
}

func (s *udpserver) serveUDP(addr string) {
	var err error
	s.udpListener, err = newUDPListener(addr)
	if err != nil {
		s.log.Fatal(err)
	}

	defer s.udpListener.Close()

	s.log.Println("serving udp", addr)

	buffer := make([]byte, 1024)

	go s.broadcast()

	for {
		_, addr, err := s.udpListener.ReadFromUDP(buffer)
		if err != nil {
			s.log.Println(err)
			continue
		}

		s.log.Println("new udp client connected", addr.String(), string(buffer))
		s.addClientAddr(addr)
	}
}

func newUDPListener(address string) (*net.UDPConn, error) {
	s, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, err
	}

	return net.ListenUDP("udp", s)
}

func (s *udpserver) broadcast() {
	for data := range s.inputDataChan {
		for _, addr := range s.clientAddrs {
			if _, err := s.udpListener.WriteToUDP(data, addr); err != nil {
				s.log.Println(err)
				s.removeClientAddr(addr)
			}
		}
	}
}

func (s *udpserver) addClientAddr(addr *net.UDPAddr) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.clientAddrs = append(s.clientAddrs, addr)
}

func (s *udpserver) removeClientAddr(addr *net.UDPAddr) {
	s.lock.Lock()
	defer s.lock.Unlock()
	for i, clientAddr := range s.clientAddrs {
		if addr == clientAddr {
			s.clientAddrs = append(s.clientAddrs[:i], s.clientAddrs[i+1:]...)
		}
	}
}
