package namegen

import (
	"github.com/jaswdr/faker"
)

type Countries struct {
	parameters NameGeneratorParameters
	takenNames map[string]int
	faker      faker.Faker
}

func NewCountries(parameters NameGeneratorParameters) *Countries {
	return &Countries{
		parameters: parameters,
		takenNames: map[string]int{},
		faker:      faker.New(),
	}
}

func (g *Countries) NewName() (string, error) {
	return newName(g.parameters, g.takenNames, func() string {
		return g.faker.Address().Country()
	})
}
