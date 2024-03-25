package searchSpaceGeneration

import (
	"fmt"
)

type nameGenerator interface {
	NewName() string
}

func (ss *SearchSpace) nameGenerator() (nameGenerator, error) {
	switch ss.NameGenerator {
	case "default":
		return newDefaultNameGenerator(), nil
	case "3letters":
		return newThreeLetterNameGenerator(), nil
	default:
		return nil, fmt.Errorf("unknown name_generator: %s", ss.NameGenerator)
	}
}

func newDefaultNameGenerator() nameGenerator {
	return newThreeLetterNameGenerator()
}
