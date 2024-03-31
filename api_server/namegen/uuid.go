package namegen

import (
	"fmt"

	"github.com/google/uuid"
)

type Uuid struct {
	parameters NameGeneratorParameters
	takenNames map[string]int
}

func NewUuid(parameters NameGeneratorParameters) *Uuid {
	return &Uuid{
		parameters: parameters,
		takenNames: map[string]int{},
	}
}

func (g *Uuid) NewName() (string, error) {
	if g.parameters.GetAllowDuplicates() {
		return g.newNameWithDuplicates(), nil
	}
	return g.newNameWithoutDuplicates()
}

func (g *Uuid) newNameWithDuplicates() string {
	originalNewName := uuid.New().String()
	finalNewName := originalNewName
	count := g.takenNames[originalNewName]
	if count > 0 {
		finalNewName = fmt.Sprintf("%s%d", originalNewName, count)
	}
	g.takenNames[originalNewName] += 1
	return finalNewName
}

func (g *Uuid) newNameWithoutDuplicates() (string, error) {
	attemptCount := 0
	for attemptCount < g.parameters.GetMaximumSampleAttempts() {
		potentialName := uuid.New().String()
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
