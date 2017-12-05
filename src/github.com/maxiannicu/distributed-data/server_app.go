package main

import (
	"github.com/maxiannicu/distributed-data/network"
	"log"
	"os"
)

func serve(channel *network.TcpChannel) {
	for ;channel.IsAlive(); {
		if bytes, err := channel.Read(); err == nil {
			log.Println(string(bytes))
		}
	}

	log.Println("Connection closed")
}

func main() {
	server, e := network.NewTcpServer(network.NewEndPoint("localhost", 31012))

	if e != nil {
		log.Panic(e)
		os.Exit(-1)
	}

	for {
		conn, err := server.AcceptConnection()

		if err != nil {
			log.Panic(err)
		} else {
			go serve(conn)
		}
	}

}
