package mediator

import (
	"github.com/maxiannicu/distributed-data/network_dto"
	"github.com/maxiannicu/distributed-data/model"
	"github.com/maxiannicu/distributed-data/network"
	"github.com/maxiannicu/distributed-data/utils"
	"strings"
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
	nodeDataResponse := network_dto.NodeDataResponse{}
	if masterEndPoint, ok := application.findMasterNode(); ok {
		channel, err := network.NewTcpChannelAsClient(*masterEndPoint)

		if err != nil {
			application.logger.Panic(err)
		}

		application.logger.Println("Sending data request to master node")
		bytes, err := network_dto.NewRequest(network_dto.GetNodeDataRequestType, "")
		if err != nil {
			application.logger.Panic(err)
		}

		channel.Write(bytes)

		if response, err := network.NextResponse(channel); err == nil {
			if err = utils.Deserealize(response.ContentType, response.Data, &nodeDataResponse); err != nil {
				application.logger.Panic(err)
			}
			application.logger.Println("Data received from master node")
		} else {
			application.logger.Panic(err)
		}

		if len(dataRequest.OrderBy) > 0 {
			application.logger.Println("Ordering data")
			var sortFunc func(a, b model.Person) int

			switch strings.ToLower(dataRequest.OrderBy) {
			case "firstname":
				sortFunc = func(a, b model.Person) int {
					return strings.Compare(a.FirstName, b.FirstName)
				}
			case "lastname":
				sortFunc = func(a, b model.Person) int {
					return strings.Compare(a.LastName, b.LastName)
				}
			case "age":
				sortFunc = func(a, b model.Person) int {
					return int(a.Age) - int(b.Age)
				}
			}

			if sortFunc != nil {
				length := nodeDataResponse.Size
				for i := 0; i < length; i++ {
					for e := i + 1; e < length; e++ {
						if sortFunc(nodeDataResponse.Data[i], nodeDataResponse.Data[e]) > 0 {
							nodeDataResponse.Data[i], nodeDataResponse.Data[e] = nodeDataResponse.Data[e], nodeDataResponse.Data[i]
						}
					}
				}
			}
		}
	} else {
		application.logger.Println("Unable to find master node")
	}

	return network_dto.NewResponse(dataRequest.Accept, network_dto.DataResponse{
		Data: nodeDataResponse.Data,
		Size: nodeDataResponse.Size,
	})
}
