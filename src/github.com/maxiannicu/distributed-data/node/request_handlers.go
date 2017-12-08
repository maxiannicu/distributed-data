package node

import (
	"github.com/maxiannicu/distributed-data/network"
	"github.com/maxiannicu/distributed-data/network_dto"
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
	if request, err := network.NextRequest(channel); err == nil {
		if request.Type == network_dto.GetNodeDataRequestType {
			if responseBytes, err := network_dto.NewResponse(application.data); err == nil {
				channel.Write(responseBytes)
			} else {
				application.logger.Panic(err)
			}
		} else {
			application.logger.Panic("This request type is not acceptable")
		}
	} else {
		if !channel.IsAlive() {
			return
		}
		application.logger.Panic(err)
	}
}
