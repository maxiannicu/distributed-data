package mediator

import (
	"github.com/maxiannicu/distributed-data/network_dto"
)

type ApplicationConfig struct {
	ListenEndPoint    network_dto.EndPoint
	DiscoveryEndPoint network_dto.EndPoint
	DiscoveryDuration int
}
