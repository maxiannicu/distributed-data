package network_dto

import "github.com/maxiannicu/distributed-data/utils"

type Response struct {
	ContentType byte
	Data        []byte
}

func NewResponse(contentType byte, data interface{}) ([]byte, error) {
	dataBytes, err := utils.Serialize(contentType, data)

	if err != nil {
		return nil, err
	}

	response := &Response{
		ContentType: contentType,
		Data:        dataBytes,
	}

	return utils.Serialize(utils.JsonFormat, response)
}
