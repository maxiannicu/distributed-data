package main

import (
	"github.com/maxiannicu/distributed-data/network"
	"log"
	"os"
)

func main()  {
	sender, err := network.NewUdpSender(network.NewEndPoint("127.0.0.1", 31012))
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}

	sender.Write([]byte("Hi"))
	sender.Write([]byte("How are you"))
}
