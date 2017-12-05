package network

import (
	"net"
	"bufio"
	"io"
)

type TcpChannel struct {
	conn       *net.TCPConn
	reader     *bufio.Reader
	writer     *bufio.Writer
	eofOccured bool
}

func NewTcpChannel(conn *net.TCPConn) *TcpChannel {
	return &TcpChannel{
		conn:   conn,
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
	}
}

func NewTcpChannelAsClient(remoteEndPoint EndPoint) (*TcpChannel, error) {
	addr, err := net.ResolveTCPAddr("tcp", remoteEndPoint.String())

	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, addr)

	if err != nil {
		return nil, err
	}

	return &TcpChannel{
		conn:   conn,
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
	}, nil
}

func (channel *TcpChannel) Read() ([]byte, error) {
	bytes, err := channel.reader.ReadBytes(messageDelimiter)

	if err != nil {
		if err == io.EOF {
			channel.eofOccured = true
		}
		return nil, err
	}

	return bytes[:len(bytes)-1], nil
}

func (channel *TcpChannel) Write(bytes []byte) {
	channel.writer.Write(append(bytes, messageDelimiter))
	channel.writer.Flush()
}

func (channel *TcpChannel) IsAlive() bool {
	return !channel.eofOccured
}

func (channel *TcpChannel) Close() {
	channel.conn.Close()
}
