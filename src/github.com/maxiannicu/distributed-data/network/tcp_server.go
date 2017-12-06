package network

import (
	"net"
)

type TcpServer struct {
	listener      *net.TCPListener
	localEndPoint EndPoint
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
		listener:      listener,
		localEndPoint: e,
	}, nil
}

func (server *TcpServer) AcceptConnection() (*TcpChannel, error) {
	conn, err := server.listener.AcceptTCP()

	if err != nil {
		return nil, err
	}

	return NewTcpChannel(conn), nil
}

func (server *TcpServer) LocalEndPoint() EndPoint {
	return server.localEndPoint
}

func (server *TcpServer) Close() {
	server.listener.Close()
}
