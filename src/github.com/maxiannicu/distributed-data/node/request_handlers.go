package node

import (
	"github.com/maxiannicu/distributed-data/network"
	"github.com/maxiannicu/distributed-data/network_dto"
	"github.com/maxiannicu/distributed-data/model"
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
	if request, err := network.NextRequest(channel); err == nil {
		if request.RequestType == network_dto.GetNodeDataRequestType {
			application.logger.Println("Received GetNodeDataRequest")
			if data, err := application.accumulateData(); err == nil {
				if responseBytes, err := network_dto.NewResponse(application.contentType, network_dto.NodeDataResponse{
					Data: data,
					Size: len(data),
				}); err == nil {
					application.logger.Println("Sending", len(data), "elements")
					channel.Write(responseBytes)
				} else {
					application.logger.Panic(err)
				}
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

func (application *Application) accumulateData() ([]model.Person, error) {
	accumulate := make([]model.Person, 0)

	if !application.inTransaction {
		application.inTransaction = true
		application.logger.Println("Added own entries total of", len(application.data))
		for _, el := range application.data {
			accumulate = append(accumulate, el)
		}

		for _, conn := range application.clients {
			if bytes, err := network_dto.NewRequest(network_dto.GetNodeDataRequestType, ""); err == nil {
				conn.Write(bytes)

				if response, err := network.NextResponse(conn); err == nil {
					dataResponse := network_dto.NodeDataResponse{}
					if err := utils.Deserealize(response.ContentType, response.Data, &dataResponse); err == nil {
						application.logger.Println("Adding received data from node with total of", dataResponse.Size)
						for _, el := range dataResponse.Data {
							accumulate = append(accumulate, el)
						}
					} else {
						application.logger.Panic(err)
					}
				} else {
					application.logger.Panic(err)
				}
			} else {
				return nil, err
			}
		}

		application.inTransaction = false
	} else {
		application.logger.Println("This node is already in transaction. Will not call connections.")
	}

	return accumulate, nil
}
