package mediator

import "github.com/maxiannicu/distributed-data/network_dto"

func getData(dataRequest network_dto.DataRequest) []byte {
	return []byte("Hello world")
}
