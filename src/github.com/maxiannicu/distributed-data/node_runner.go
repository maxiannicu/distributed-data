package main

import (
	"github.com/maxiannicu/distributed-data/utils"
	"log"
	"github.com/maxiannicu/distributed-data/node"
	"time"
)

var logger = utils.NewLogger("node_runner")

func main() {
	configs, err := getNodeConfigs()
	if err != nil {
		log.Panic(err)
	}

	nodes := startNodes(configs)
	connectNodes(configs, nodes)

	time.Sleep(time.Hour * 24)
}

func connectNodes(configs []node.ApplicationConfig, nodes map[string]*node.Application) {
	for _, config := range configs {
		for _, conn := range config.Connections {
			connectedNode := nodes[conn]
			err := nodes[config.Identificator].ConnectTo(connectedNode.LocalEndPoint())
			if err != nil {
				logger.Panic(err)
			}
		}
	}
}

func startNodes(configs []node.ApplicationConfig) map[string]*node.Application {
	nodes := make(map[string]*node.Application)
	for _, config := range configs {
		if application, err := node.NewApplication(config); err != nil {
			logger.Println("Unable to start", config.Identificator, ".Reason :", err)
		} else {
			nodes[config.Identificator] = application
			go application.Loop()
		}
	}
	return nodes
}

func getNodeConfigs() ([]node.ApplicationConfig, error) {
	fileNames, err := utils.GetFileNamesByPattern("^node.[\\d]+.json$")
	if err != nil {
		log.Panic(err)
	}
	configs := make([]node.ApplicationConfig, 0)

	for _, val := range fileNames {
		config := node.ApplicationConfig{}
		err := utils.GetConfig(val, &config)

		if err != nil {
			return nil, err
		}

		configs = append(configs, config)
	}

	logger.Println("Loaded",len(configs),"node configurations")

	return configs, nil
}
