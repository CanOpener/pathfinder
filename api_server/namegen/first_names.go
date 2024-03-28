package namegen

import (
	"github.com/jaswdr/faker"
)

type FirstNames struct {
	faker faker.Faker
}

func NewFirstNames() FirstNames {
	return FirstNames{
		faker: faker.New(),
	}
}

func (g FirstNames) NewName() string {
	return g.faker.Person().FirstName()
}
