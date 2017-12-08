package node

import (
	"github.com/maxiannicu/distributed-data/network"
	"github.com/maxiannicu/distributed-data/network_dto"
	"github.com/maxiannicu/distributed-data/utils"
)

func (application *Application) handleTcpRequests() {
	for {
		channel, err := application.server.AcceptConnection()

		if err != nil {
			application.logger.Panic(err)
		}

		go application.handleTcpChannel(channel)
	}
}

func (application *Application) handleTcpChannel(channel *network.TcpChannel) {
	bytes, err := channel.Read()

	if !channel.IsAlive() {
		return
	} else if err != nil {
		application.logger.Panic(err)
	}

	request := network_dto.Request{}

	err = utils.Deserealize(utils.JsonFormat, bytes, &request)
	if err != nil {
		application.logger.Panic(err)
	}

	if request.Type == network_dto.GetNodeDataRequestType {
		responseBytes, err := network_dto.NewResponse(application.data)

		if err != nil {
			application.logger.Panic(err)
		} else {
			channel.Write(responseBytes)
		}
	} else {
		application.logger.Panic("This request type is not acceptable")
	}
}
