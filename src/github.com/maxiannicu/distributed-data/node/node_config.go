package node

import "github.com/maxiannicu/distributed-data/data"

type ApplicationConfig struct {
	Identificator string
	Connections []string
	Data []data.Person
}
