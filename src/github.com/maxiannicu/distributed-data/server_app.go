package main

import (
	"github.com/maxiannicu/distributed-data/network"
	"log"
	"os"
	"time"
)

func main() {
	go CreateNewListener()
	//go CreateNewListener()

	time.Sleep(time.Second * 10)
}
func CreateNewListener() {
	listener, err := network.NewUdpListener(network.NewEndPoint("127.0.0.1", 31012))
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}
	for {
		bytes, err := listener.Read()

		if err != nil {
			log.Panic(err)
		} else {
			log.Println(string(bytes))
		}
	}
}
