package network

import "fmt"

type EndPoint struct {
	host string
	port int
}

func (e EndPoint) String() string {
	return fmt.Sprintf("%s:%d", e.host, e.port)
}

func NewEndPoint(host string, port int) EndPoint {
	return EndPoint{
		host:host,
		port:port,
	}
}