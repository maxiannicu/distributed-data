package mediator

import (
	"github.com/maxiannicu/distributed-data/network"
	"log"
	"github.com/maxiannicu/distributed-data/utils"
	"time"
)

type Application struct {
	listeningServer      *network.TcpServer
	discoveryUdpSender   *network.UdpSender
	discoveryUdpListener *network.UdpListener
	discoveryDuration    time.Duration
	logger               *log.Logger
}

func NewApplication(config ApplicationConfig) (*Application, error) {
	logger := utils.NewLogger("mediator")
	logger.Println("Starting TCP server")
	server, err := network.NewTcpServerWithEndpoint(config.ListenEndPoint)

	if err != nil {
		return nil, err
	}
	logger.Println("TCP server started on", server.LocalEndPoint())

	logger.Println("Starting UDP sender")
	sender, err := network.NewUdpSender(config.DiscoveryEndPoint)
	if err != nil {
		return nil, err
	}
	logger.Println("UDP sender started")

	logger.Println("Starting UDP listener")
	listener, err := network.NewUdpListener()
	if err != nil {
		return nil, err
	}
	logger.Println("UDP listener started on", listener.LocalEndPoint())

	return &Application{
		listeningServer:      server,
		discoveryUdpSender:   sender,
		discoveryUdpListener: listener,
		discoveryDuration:    time.Millisecond * time.Duration(config.DiscoveryDuration),
		logger:               logger,
	}, nil
}

func (application *Application) Loop() {
	for {
		channel, err := application.listeningServer.AcceptConnection()

		if err != nil {
			application.logger.Panic(err)
		} else {
			go application.handleClient(channel)
		}
	}
}
