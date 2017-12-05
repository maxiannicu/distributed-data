package data

import (
	"math/rand"
	"time"
)

type PersonRepository struct {
	data []Person
}

var firstNames = []string{
	"Alex",
	"Greg",
	"Michael",
	"Andrew",
	"Joseph",
}

var flastNames = []string{
	"Young",
	"Cole",
	"Milner",
	"O'Connor",
	"Achtenberg",
}

func NewPersonRepository() *PersonRepository {
	rand.Seed(time.Now().UTC().UnixNano())
	cnt := rand.Intn(10)
	people := make([]Person, cnt)

	for i := 0; i < cnt; i++ {
		people[i] = Person{
			FirstName:firstNames[rand.Int()%len(firstNames)],
			LastName:flastNames[rand.Int()%len(flastNames)],
			Age: byte(20 + rand.Intn(40)),
		}
	}

	return &PersonRepository{
		data:people,
	}
}

func (repository *PersonRepository) Get() []Person {
	return repository.data
}