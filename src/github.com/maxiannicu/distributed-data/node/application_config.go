package node

import (
	"github.com/maxiannicu/distributed-data/model"
	"github.com/maxiannicu/distributed-data/network"
)

type ApplicationConfig struct {
	Identificator     string
	Connections       []string
	Data              []model.Person
	DiscoveryEndPoint network.EndPoint
}
