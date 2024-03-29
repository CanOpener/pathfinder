package ssgen

import (
	"fmt"

	"github.com/canopener/pathfinder/api_server/namegen"
)

type nameGenerator interface {
	NewName() string
}

func newNameGenerator(generatorId string) (nameGenerator, error) {
	switch generatorId {
	case "three_letters":
		return namegen.NewThreeLetters(), nil
	case "cities":
		return namegen.NewCities(), nil
	case "countries":
		return namegen.NewCountries(), nil
	case "first_names":
		return namegen.NewFirstNames(), nil
	default:
		return nil, fmt.Errorf("unknown name_generator_id: %s", generatorId)
	}
}
