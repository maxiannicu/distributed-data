package node

import (
	"github.com/maxiannicu/distributed-data/data"
	"github.com/maxiannicu/distributed-data/network"
)

type ApplicationConfig struct {
	Identificator     string
	Connections       []string
	Data              []data.Person
	DiscoveryEndPoint network.EndPoint
}
