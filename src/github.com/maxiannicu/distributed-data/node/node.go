package node

import (
	"github.com/maxiannicu/distributed-data/network"
	"github.com/maxiannicu/distributed-data/data"
	"log"
	"github.com/maxiannicu/distributed-data/utils"
)

type Application struct {
	server  *network.TcpServer
	clients []*network.TcpChannel
	data    []data.Person
	logger *log.Logger
}

func NewApplication(config ApplicationConfig) (*Application, error) {
	logger := utils.NewLogger(config.Identificator)
	logger.Println("Starting tcp server")
	server, err := network.NewTcpServer()

	if err != nil {
		return nil, err
	}
	logger.Println("Tcp server started",server.LocalEndPoint())

	return &Application{
		server:  server,
		clients: make([]*network.TcpChannel, 0),
		data:    config.Data,
		logger:  logger,
	}, nil
}

func (application *Application) ConnectTo(remoteEndPoint network.EndPoint) error {
	application.logger.Println("Connecting to",remoteEndPoint)
	channel, err := network.NewTcpChannelAsClient(remoteEndPoint)
	if err != nil {
		return err
	}

	application.logger.Println("Connected",remoteEndPoint)
	application.clients = append(application.clients, channel)

	return nil
}

func (application *Application) LocalEndPoint() network.EndPoint {
	return application.server.LocalEndPoint()
}

func (application *Application) Loop() {
	application.logger.Println("Looping")
}