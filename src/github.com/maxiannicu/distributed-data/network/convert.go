package network

import (
	"net"
	"fmt"
)

func toEndPoint(address net.Addr) EndPoint {
	endPoint := EndPoint{}
	fmt.Sscanf(address.String(),"%s:%d", &endPoint.Host, &endPoint.Port)

	return endPoint
}
