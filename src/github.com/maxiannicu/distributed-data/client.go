package main

import (
	"github.com/maxiannicu/distributed-data/network"
	"log"
	"github.com/maxiannicu/distributed-data/network_dto"
	"github.com/maxiannicu/distributed-data/utils"
	"fmt"
)

func main() {
	channel, err := network.NewTcpChannelAsClient(network_dto.NewEndPoint("127.0.0.1", 31012))

	if err != nil {
		log.Panic(err)
	}

	bytes, err := network_dto.NewRequest(network_dto.GetDataRequestType, network_dto.DataRequest{
		OrderBy: "FirstName",
		Accept:  utils.XMLFormat,
	})
	if err != nil {
		log.Panic(err)
	}
	channel.Write(bytes)

	bytes, err = channel.Read()

	if err != nil {
		log.Panic(err)
	}

	response := network_dto.Response{}
	utils.Deserealize(utils.JsonFormat, bytes, &response)

	fmt.Println(string(response.Data))

	channel.Close()
}
