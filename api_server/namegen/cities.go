package namegen

import (
	"github.com/jaswdr/faker"
)

type Cities struct {
	faker faker.Faker
}

func NewCities() Cities {
	return Cities{
		faker: faker.New(),
	}
}

func (g Cities) NewName() string {
	return g.faker.Address().City()
}
