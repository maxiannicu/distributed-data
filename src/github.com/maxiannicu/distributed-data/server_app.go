package main

import (
	"github.com/maxiannicu/distributed-data/network"
	"log"
	"os"
	"github.com/maxiannicu/distributed-data/data"
	"github.com/maxiannicu/distributed-data/utils"
	"time"
)

func main()  {
	repository := data.NewPersonRepository()

	sender, err := network.NewUdpSender(network.NewEndPoint("224.0.0.1", 31012))
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}
	defer sender.Close()

	for {
		bytes, err := utils.Serialize(utils.JsonFormat, repository.Get())
		if err != nil {
			log.Panic(err)
		} else {
			sender.Write(bytes)
		}
		time.Sleep(time.Second * 1)
	}
}

