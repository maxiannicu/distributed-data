package mediator

import (
	"github.com/maxiannicu/distributed-data/network_dto"
	"time"
	"io"
	"github.com/maxiannicu/distributed-data/utils"
)

func (application *Application) findMasterNode() {
	bytes, err := network_dto.CreateRequest(network_dto.DiscoveryRequestType, network_dto.DiscoveryRequest{
		ResponseEndPoint: application.discoveryUdpListener.LocalEndPoint(),
	})

	if err != nil {
		application.logger.Fatal(err)
	}

	application.logger.Println("Sending discovery command")
	application.discoveryUdpSender.Write(bytes)
	application.logger.Println("Waiting")
	time.Sleep(application.discoveryDuration)

	responses := application.getDiscoveredNodes()

	application.logger.Println("Found", len(responses), "nodes")
}
func (application *Application) getDiscoveredNodes() []network_dto.DiscoveryResponse {
	responses := make([]network_dto.DiscoveryResponse, 0)
	for ; application.discoveryUdpListener.HasBytesAvailable(); {
		bytes, err := application.discoveryUdpListener.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			application.logger.Fatal(err)
		} else {
			response := network_dto.DiscoveryResponse{}
			utils.Deserealize(utils.JsonFormat, bytes, &response)
			responses = append(responses, response)
		}
	}
	return responses
}
