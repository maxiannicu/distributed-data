package mediator

import (
	"github.com/maxiannicu/distributed-data/network"
	"log"
	"github.com/maxiannicu/distributed-data/utils"
	"github.com/maxiannicu/distributed-data/network_dto"
	"time"
)

type Application struct {
	listeningServer    *network.TcpServer
	discoveryUdpSender *network.UdpSender
	discoveryUdpListener *network.UdpListener
	discoveryDuration time.Duration
}

func NewApplication(config ApplicationConfig) (*Application, error) {
	server, err := network.NewTcpServer(config.ListenEndPoint)

	if err != nil {
		return nil, err
	}

	sender, err := network.NewUdpSender(config.MulticastEndPoint)
	if err != nil {
		return nil, err
	}


	listener, err := network.NewUdpListener(config.MulticastEndPoint)
	if err != nil {
		return nil, err
	}

	return &Application{
		listeningServer:    server,
		discoveryUdpSender: sender,
		discoveryUdpListener: listener,
		discoveryDuration: time.Millisecond * time.Duration(config.DiscoveryDuration),
	}, nil
}

func (application *Application) Listen() {
	for {
		channel, err := application.listeningServer.AcceptConnection()

		if err != nil {
			log.Panic(err)
		} else {
			go application.handleClient(channel)
		}
	}
}

func (application *Application) handleClient(channel *network.TcpChannel) {
	log.Println("Connection open with ", channel.RemoteEndPoint())

	for ; channel.IsAlive(); {
		bytes, err := channel.Read()

		if err != nil {
			log.Println("Error reading on ", channel.RemoteEndPoint(), ". Error : ", err)
		} else {
			request := network_dto.Request{}
			err := utils.Deserealize(utils.JsonFormat, bytes, &request)
			if err != nil {
				log.Panic(err)
			} else {
				dataRequest := network_dto.DataRequest{}
				err := utils.Deserealize(utils.JsonFormat, request.Data, &dataRequest)
				if err != nil {
					log.Panic(err)
				} else {
					application.findMasterNode()
				}
			}
		}
	}
}