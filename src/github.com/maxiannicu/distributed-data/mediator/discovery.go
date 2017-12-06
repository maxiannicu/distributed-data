package mediator

import (
	"github.com/maxiannicu/distributed-data/network_dto"
	"log"
	"time"
	"io"
	"github.com/maxiannicu/distributed-data/utils"
)

func (application *Application) findMasterNode() {
	bytes, err := network_dto.CreateRequest(network_dto.DiscoveryRequestType, network_dto.DiscoveryRequest{
		ResponseEndPoint: application.discoveryUdpListener.LocalEndPoint(),
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Sending discovery command")
	application.discoveryUdpSender.Write(bytes)
	log.Println("Waiting")
	time.Sleep(application.discoveryDuration)

	responses := application.getDiscoveredNodes()

	log.Println("Found", len(responses), "nodes")
}
func (application *Application) getDiscoveredNodes() []network_dto.DiscoveryResponse {
	responses := make([]network_dto.DiscoveryResponse, 0)
	for {
		bytes, err := application.discoveryUdpListener.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		} else {
			response := network_dto.DiscoveryResponse{}
			utils.Deserealize(utils.JsonFormat, bytes, &response)
			responses = append(responses, response)
		}
	}
	return responses
}
