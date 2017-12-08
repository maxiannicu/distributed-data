package mediator

import (
	"github.com/maxiannicu/distributed-data/network_dto"
	"github.com/maxiannicu/distributed-data/model"
	"github.com/maxiannicu/distributed-data/network"
	"github.com/maxiannicu/distributed-data/utils"
)

func (application *Application) handleRequest(dataRequest network_dto.DataRequest) ([]byte, error) {
	responseData := make([]model.Person, 0)
	if masterEndPoint, ok := application.findMasterNode(); ok {
		channel, err := network.NewTcpChannelAsClient(*masterEndPoint)

		if err != nil {
			application.logger.Panic(err)
		}

		bytes, err := network_dto.CreateRequest(network_dto.GetNodeDataRequestType, "")
		if err != nil {
			application.logger.Panic(err)
		}

		channel.Write(bytes)
		responseBytes, err := channel.Read()

		if err != nil {
			application.logger.Panic(err)
		}

		response := network_dto.Response{}
		err = utils.Deserealize(utils.JsonFormat, responseBytes, &response)

		if err != nil {
			application.logger.Panic(err)
		}

		err = utils.Deserealize(utils.JsonFormat, response.Data, &responseData)
		if err != nil {
			application.logger.Panic(err)
		}
	}

	return network_dto.NewResponse(responseData)
}
