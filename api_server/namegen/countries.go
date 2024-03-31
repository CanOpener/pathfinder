package namegen

import (
	"fmt"

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
	if g.parameters.GetAllowDuplicates() {
		return g.newNameWithDuplicates(), nil
	}
	return g.newNameWithoutDuplicates()
}

func (g *Countries) newNameWithDuplicates() string {
	originalNewName := g.faker.Address().Country()
	finalNewName := originalNewName
	count := g.takenNames[originalNewName]
	if count > 0 {
		finalNewName = fmt.Sprintf("%s%d", originalNewName, count)
	}
	g.takenNames[originalNewName] += 1
	return finalNewName
}

func (g *Countries) newNameWithoutDuplicates() (string, error) {
	attemptCount := 0
	for attemptCount < g.parameters.GetMaximumSampleAttempts() {
		potentialName := g.faker.Address().Country()
		if _, ok := g.takenNames[potentialName]; ok {
			attemptCount += 1
			continue
		}

		g.takenNames[potentialName] = 1
		return potentialName, nil
	}

	return "", fmt.Errorf("failed to find an unused name after %d attempts. total taken names: %d",
		attemptCount, len(g.takenNames))
}
