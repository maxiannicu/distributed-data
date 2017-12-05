package network

import (
	"net"
	"bufio"
)

type MulticastUdpListener struct {
	conn   *net.UDPConn
	reader *bufio.Reader
}

func NewMulticastUdpListener(listenEndPoint EndPoint) (*MulticastUdpListener, error) {
	addr, err := net.ResolveUDPAddr("udp", listenEndPoint.String())
	if err != nil {
		return nil, err
	}

	listen, err := net.ListenMulticastUDP("udp", nil, addr)

	if err != nil {
		return nil, err
	}

	return &MulticastUdpListener{
		conn:   listen,
		reader: bufio.NewReader(listen),
	}, nil
}

func (listener *MulticastUdpListener) Read() ([]byte, error) {
	bytes, err := listener.reader.ReadBytes(messageDelimiter)

	if err != nil {
		return nil, err
	}

	return bytes[:len(bytes)-1], nil
}

func (listener *MulticastUdpListener) Close() {
	listener.conn.Close()
}
