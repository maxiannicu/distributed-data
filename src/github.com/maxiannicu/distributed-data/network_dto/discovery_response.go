package network_dto

import "github.com/maxiannicu/distributed-data/network"

type DiscoveryResponse struct {
	DataSize int
	ConnectionEndPoint network.EndPoint
}