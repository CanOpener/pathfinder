package namegen

import (
	"fmt"

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
	if g.parameters.GetAllowDuplicates() {
		return g.newNameWithDuplicates(), nil
	}
	return g.newNameWithoutDuplicates()
}

func (g *FirstNames) newNameWithDuplicates() string {
	originalNewName := g.faker.Person().FirstName()
	finalNewName := originalNewName
	count := g.takenNames[originalNewName]
	if count > 0 {
		finalNewName = fmt.Sprintf("%s%d", originalNewName, count)
	}
	g.takenNames[originalNewName] += 1
	return finalNewName
}

func (g *FirstNames) newNameWithoutDuplicates() (string, error) {
	attemptCount := 0
	for attemptCount < g.parameters.GetMaximumSampleAttempts() {
		potentialName := g.faker.Person().FirstName()
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
