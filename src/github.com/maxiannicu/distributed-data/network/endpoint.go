package network

import "fmt"

type EndPoint struct {
	Host string
	Port int
}

func (e EndPoint) String() string {
	return fmt.Sprintf("%s:%d", e.Host, e.Port)
}

func NewEndPoint(host string, port int) EndPoint {
	return EndPoint{
		Host: host,
		Port: port,
	}
}