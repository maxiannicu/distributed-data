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
		if request.Type == network_dto.GetNodeDataRequestType {
			application.logger.Println("Received GetNodeDataRequest")
			if data, err := application.accumulateData(); err == nil {
				if responseBytes, err := network_dto.NewResponse(data); err == nil {
					application.logger.Println("Sending response back")
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
		for _, el := range application.data {
			accumulate = append(accumulate, el)
		}

		for _, conn := range application.clients {
			if bytes, err := network_dto.NewRequest(network_dto.GetNodeDataRequestType, ""); err == nil {
				conn.Write(bytes)

				if response, err := network.NextResponse(conn); err == nil {
					clientData := make([]model.Person, 0)
					if err := utils.Deserealize(utils.JsonFormat, response.Data, &clientData); err == nil {
						for _, el := range clientData {
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
