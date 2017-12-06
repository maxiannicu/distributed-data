package network

import (
	"net"
	"bufio"
)

type UdpListener struct {
	conn   *net.UDPConn
	reader *bufio.Reader
}

func NewUdpListener(listenEndPoint EndPoint) (*UdpListener, error) {
	addr, err := net.ResolveUDPAddr("udp", listenEndPoint.String())
	if err != nil {
		return nil, err
	}

	listen, err := net.ListenUDP("udp", addr)

	if err != nil {
		return nil, err
	}

	return &UdpListener{
		conn:   listen,
		reader: bufio.NewReader(listen),
	}, nil
}

func (listener *UdpListener) Read() ([]byte, error) {
	bytes, err := listener.reader.ReadBytes(messageDelimiter)

	if err != nil {
		return nil, err
	}

	return bytes[:len(bytes)-1], nil
}

func (listener *UdpListener) LocalEndPoint() EndPoint {
	return toEndPoint(listener.conn.LocalAddr())
}

func (listener *UdpListener) Close() {
	listener.conn.Close()
}
