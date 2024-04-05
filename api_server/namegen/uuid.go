package namegen

import (
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
	return newName(g.parameters, g.takenNames, func() string {
		return uuid.New().String()
	})
}
