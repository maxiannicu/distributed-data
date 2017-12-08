package network

import (
	"net"
	"bufio"
	"time"
	"github.com/maxiannicu/distributed-data/network_dto"
)

type UdpListener struct {
	conn   *net.UDPConn
	reader *bufio.Reader
}


func createUdpListener(localAddr *net.UDPAddr) (*UdpListener, error) {
	listen, err := net.ListenUDP("udp", localAddr)

	if err != nil {
		return nil, err
	}

	return &UdpListener{
		conn:   listen,
		reader: bufio.NewReader(listen),
	}, nil
}

func NewUdpListenerWithEndpoint(listenEndPoint network_dto.EndPoint) (*UdpListener, error) {
	addr, err := net.ResolveUDPAddr("udp", listenEndPoint.String())
	if err != nil {
		return nil, err
	}

	return createUdpListener(addr)
}

func NewUdpListener() (*UdpListener, error) {
	return createUdpListener(nil)
}

func (listener *UdpListener) Read() ([]byte, error) {
	bytes, err := listener.reader.ReadBytes(messageDelimiter)

	if err != nil {
		return nil, err
	}

	return bytes[:len(bytes)-1], nil
}

func (listener *UdpListener) LocalEndPoint() network_dto.EndPoint {
	return network_dto.NewEndPointFromAddr(listener.conn.LocalAddr())
}

func (listener *UdpListener) SetReadTimeOut(time time.Time) {
	listener.conn.SetReadDeadline(time)
}


func (listener *UdpListener) Close() {
	listener.conn.Close()
}
