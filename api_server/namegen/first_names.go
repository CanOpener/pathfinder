package namegen

import (
	"github.com/jaswdr/faker"
)

type FirstNames struct {
	parameters NameGeneratorParameters
	takenNames map[string]int
	faker      faker.Faker
}

func NewFirstNames(parameters NameGeneratorParameters) *FirstNames {
	return &FirstNames{
		parameters: parameters,
		takenNames: map[string]int{},
		faker:      faker.New(),
	}
}

func (g *FirstNames) NewName() (string, error) {
	return newName(g.parameters, g.takenNames, func() string {
		return g.faker.Person().FirstName()
	})
}
