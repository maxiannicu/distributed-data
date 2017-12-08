package network_dto

import (
	"github.com/maxiannicu/distributed-data/utils"
)

type Request struct {
	Type byte
	Data []byte
}

func NewRequest(requestType byte, data interface{}) ([]byte, error) {
	serializedData, err := utils.Serialize(utils.JsonFormat, data)
	if err != nil {
		return nil, err
	}

	return utils.Serialize(utils.JsonFormat, Request{
		Type: requestType,
		Data: serializedData,
	})
}
