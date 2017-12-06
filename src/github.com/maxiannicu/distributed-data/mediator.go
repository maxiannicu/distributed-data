package main

import (
	"github.com/maxiannicu/distributed-data/mediator"
	"github.com/maxiannicu/distributed-data/utils"
	"log"
)

func main() {
	config := mediator.ApplicationConfig{}
	err := utils.GetConfig("mediator.json", &config)
	if err != nil {
	    log.Panic(err)
	}
	app, err := mediator.NewApplication(config)
	if err != nil {
	    log.Panic(err)
	}

	app.Listen()
}