package main

import (
	"github.com/maxiannicu/distributed-data/network"
	"log"
	"os"
	"time"
	"github.com/maxiannicu/distributed-data/data"
	"github.com/maxiannicu/distributed-data/utils"
)

func main() {
	go CreateNewListener()
	go CreateNewListener()

	time.Sleep(time.Second * 10)
}
func CreateNewListener() {
	listener, err := network.NewMulticastUdpListener(network.NewEndPoint("224.0.0.1", 31012))
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}
	for {
		bytes, err := listener.Read()


		if err != nil {
			log.Panic(err)
		} else {
			var list = make([]data.Person, 0)
			utils.Deserealize(utils.JsonFormat, bytes, &list)

			log.Println("Received",len(list), "persons")
		}
	}
}