package network

import (
	"net"
	"strings"
	"strconv"
)

func toEndPoint(address net.Addr) EndPoint {
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
