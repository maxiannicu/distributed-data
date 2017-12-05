package network

import (
	"net"
)

type TcpServer struct {
	listener *net.TCPListener
}

func NewTcpServer(e EndPoint) (*TcpServer, error) {
	addr, err := net.ResolveTCPAddr("tcp", e.String())
	if err != nil {
		return nil, err
	}
	listener, err := net.ListenTCP("tcp", addr)

	if err != nil {
		return nil, err
	}

	return &TcpServer{
		listener: listener,
	}, nil
}

func (server *TcpServer) Close() {
	server.Close()
}

func (server *TcpServer) AcceptConnection() (*TcpChannel,error) {
	conn, err := server.listener.AcceptTCP()

	if err != nil {
		return nil, err
	}

	return NewTcpChannel(conn), nil
}
