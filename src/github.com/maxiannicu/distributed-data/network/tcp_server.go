package network

import (
	"net"
)

type TcpServer struct {
	listener *net.TCPListener
}

func createTcpServer(addr *net.TCPAddr) (*TcpServer, error) {
	listener, err := net.ListenTCP("tcp", addr)

	if err != nil {
		return nil, err
	}

	return &TcpServer{
		listener: listener,
	}, nil
}


func NewTcpServerWithEndpoint(e EndPoint) (*TcpServer, error) {
	addr, err := net.ResolveTCPAddr("tcp", e.String())
	if err != nil {
		return nil, err
	}

	return createTcpServer(addr)
}

func NewTcpServer() (*TcpServer, error) {
	return createTcpServer(nil)
}

func (server *TcpServer) AcceptConnection() (*TcpChannel, error) {
	conn, err := server.listener.AcceptTCP()

	if err != nil {
		return nil, err
	}

	return NewTcpChannel(conn), nil
}

func (server *TcpServer) LocalEndPoint() EndPoint {
	return toEndPoint(server.listener.Addr())
}

func (server *TcpServer) Close() {
	server.listener.Close()
}
