package network

import (
	"github.com/maxiannicu/distributed-data/network_dto"
	"github.com/maxiannicu/distributed-data/utils"
)

type DataChannel interface {
	Read() ([]byte, error)
	Write(bytes []byte)
}

func NextResponse(channel DataChannel) (*network_dto.Response, error) {
	if bytes, err := channel.Read(); err == nil {
		response := &network_dto.Response{}

		if err := utils.Deserealize(utils.JsonFormat, bytes, response); err == nil {
			return response, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}


func NextRequest(channel DataChannel) (*network_dto.Request, error) {
	if bytes, err := channel.Read(); err == nil {
		response := &network_dto.Request{}

		if err := utils.Deserealize(utils.JsonFormat, bytes, response); err == nil {
			return response, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
