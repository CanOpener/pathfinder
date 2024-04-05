package namegen

import (
	"github.com/jaswdr/faker"
)

type Cities struct {
	parameters NameGeneratorParameters
	takenNames map[string]int
	faker      faker.Faker
}

func NewCities(parameters NameGeneratorParameters) *Cities {
	return &Cities{
		parameters: parameters,
		takenNames: map[string]int{},
		faker:      faker.New(),
	}
}

func (g *Cities) NewName() (string, error) {
	return newName(g.parameters, g.takenNames, func() string {
		return g.faker.Address().City()
	})
}
