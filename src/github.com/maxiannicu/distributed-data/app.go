package main

import (
	"github.com/maxiannicu/distributed-data/network"
	"log"
	"time"
)

func main()  {
	endPoint := network.NewEndPoint("127.0.0.1", 31012)

	client, err := network.NewTcpChannelAsClient(endPoint)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Sending messages")
	client.Write([]byte("Hello"))
	client.Write([]byte("How are you"))

	time.Sleep(1 * time.Second)

	client.Close()
}
