package network

import (
	"net"
	"bufio"
)

type UdpSender struct {
	conn   *net.UDPConn
	writer *bufio.Writer
}

func NewUdpSender(remoteEndPoint EndPoint) (*UdpSender, error) {
	addr, err := net.ResolveUDPAddr("udp", remoteEndPoint.String())
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, err
	}

	return &UdpSender{
		conn:   conn,
		writer: bufio.NewWriter(conn),
	}, nil
}

func (sender *UdpSender) Write(bytes []byte) {
	sender.writer.Write(append(bytes, messageDelimiter))
	sender.writer.Flush()
}

func (sender *UdpSender) Close() {
	sender.conn.Close()
}
