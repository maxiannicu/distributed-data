package network_dto

import (
	"fmt"
	"net"
	"strings"
	"strconv"
)

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

func NewEndPointFromAddr(address net.Addr) EndPoint {
	split := strings.Split(address.String(), ":")
	untilPort := len(split) - 1
	port, _ := strconv.Atoi(split[untilPort])

	return EndPoint{
		Host: getHost(strings.Join(split[:untilPort], "")),
		Port: port,
	}
}

func getHost(host string) string {
	ipv4 := net.ParseIP(host).To4()
	ipAsString := ipv4.String()
	if ipv4 == nil {
		// trick for avoiding IP problem when sending details
		ipAsString = "127.0.0.1"
	}
	return ipAsString
}
