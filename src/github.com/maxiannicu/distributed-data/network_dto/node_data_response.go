package network_dto

import "github.com/maxiannicu/distributed-data/model"

type NodeDataResponse struct {
	Size int
	Data []model.Person
}
