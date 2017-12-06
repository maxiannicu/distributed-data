package main

import (
	"github.com/maxiannicu/distributed-data/network"
	"log"
)

func main()  {
	listener, err := network.NewMulticastUdpListener(network.NewEndPoint("224.0.0.1", 31013))

	if err != nil {
	    log.Panic(err)
	}

	for {
		bytes, err := listener.Read()

		if err != nil {
			log.Panic(err)
		} else {
			log.Panic(string(bytes))
		}
	}
}

