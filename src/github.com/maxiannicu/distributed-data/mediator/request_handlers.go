package mediator

import (
	"github.com/maxiannicu/distributed-data/network_dto"
	"github.com/maxiannicu/distributed-data/model"
	"github.com/maxiannicu/distributed-data/network"
	"github.com/maxiannicu/distributed-data/utils"
)

func (application *Application) handleClient(channel *network.TcpChannel) {
	application.logger.Println("Connection open with ", channel.RemoteEndPoint())

	for {
		if request, err := network.NextRequest(channel); err == nil {
			dataRequest := network_dto.DataRequest{}
			err := utils.Deserealize(utils.JsonFormat, request.Data, &dataRequest)
			if err != nil {
				application.logger.Panic(err)
			} else {
				responseBytes, err := application.handleRequest(dataRequest)
				if err != nil {
					application.logger.Panic(err)
				}

				channel.Write(responseBytes)
			}
		} else {
			if !channel.IsAlive() {
				return
			}

			application.logger.Panic(err)
		}
	}
}

func (application *Application) handleRequest(dataRequest network_dto.DataRequest) ([]byte, error) {
	responseData := make([]model.Person, 0)
	if masterEndPoint, ok := application.findMasterNode(); ok {
		channel, err := network.NewTcpChannelAsClient(*masterEndPoint)

		if err != nil {
			application.logger.Panic(err)
		}

		bytes, err := network_dto.NewRequest(network_dto.GetNodeDataRequestType, "")
		if err != nil {
			application.logger.Panic(err)
		}

		channel.Write(bytes)

		if response, err := network.NextResponse(channel); err == nil {
			if err = utils.Deserealize(response.ContentType, response.Data, &responseData); err != nil {
				application.logger.Panic(err)
			}
		} else {
			application.logger.Panic(err)
		}
	}

	return network_dto.NewResponse(dataRequest.Accept, responseData)
}
