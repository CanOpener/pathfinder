package namegen

import (
	"github.com/jaswdr/faker"
)

type Countries struct {
	faker faker.Faker
}

func NewCountries() Countries {
	return Countries{
		faker: faker.New(),
	}
}

func (g Countries) NewName() string {
	return g.faker.Address().Country()
}
