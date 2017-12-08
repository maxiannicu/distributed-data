package node

import (
	"github.com/maxiannicu/distributed-data/network_dto"
	"github.com/maxiannicu/distributed-data/utils"
	"github.com/maxiannicu/distributed-data/network"
)

func (application *Application) listenDiscovery() {
	application.logger.Println("Listening for discovery")
	for {
		bytes, err := application.discoveryListener.Read()
		if err != nil {
			application.logger.Panic(err)
		}

		request := network_dto.Request{}
		err = utils.Deserealize(utils.JsonFormat, bytes, &request)

		if err != nil {
			application.logger.Panic(err)
		}

		if request.Type != network_dto.DiscoveryRequestType {
			application.logger.Panic("Not valid request for discovery")
		}
		application.logger.Println("Discovery request received")

		discoveryRequest := network_dto.DiscoveryRequest{}
		err = utils.Deserealize(utils.JsonFormat, request.Data, &discoveryRequest)

		if err != nil {
			application.logger.Panic(err)
		}

		sender, err := network.NewUdpSender(discoveryRequest.ResponseEndPoint)

		if err != nil {
			application.logger.Panic(err)
		}

		response := network_dto.DiscoveryResponse{
			ConnectionEndPoint: application.server.LocalEndPoint(),
			DataSize:           len(application.data),
		}

		serialize, err := utils.Serialize(utils.JsonFormat, response)

		if err != nil {
			application.logger.Panic(err)
		}
		sender.Write(serialize)
		application.logger.Println("Discovery response sent back to", discoveryRequest.ResponseEndPoint)
	}
}
