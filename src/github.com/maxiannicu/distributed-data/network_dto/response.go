package network_dto

import "github.com/maxiannicu/distributed-data/utils"

type Response struct {
	Data []byte
}

func NewResponse(data interface{}) ([]byte, error) {
	dataBytes, err := utils.Serialize(utils.JsonFormat, data)

	if err != nil {
		return nil, err
	}

	response := &Response{
		Data: dataBytes,
	}

	return utils.Serialize(utils.JsonFormat, response)
}
