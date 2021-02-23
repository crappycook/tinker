package server

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
)

type TcpServer struct {
	IP         string
	Port       int
	Networking string
}

func (s *TcpServer) Start() {
	addr, err := net.ResolveTCPAddr(s.Networking, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		return
	}

	listener, err := net.ListenTCP(s.Networking, addr)
	if err != nil {
		return
	}
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Printf("Accept err: %v ", err)
			continue
		}
		go handleConn(conn) // todo
	}
}

func handleConn(conn *net.TCPConn) {
	io.Copy(ioutil.Discard, conn)
}
