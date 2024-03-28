package ssgen

import (
	"fmt"

	"github.com/canopener/pathfinder/api_server/namegen"
)

type nameGenerator interface {
	NewName() string
}

func newNameGenerator(generatorId string) (nameGenerator, string, error) {
	switch generatorId {
	case "default":
		return namegen.NewThreeLetters(), "three_letters", nil
	case "three_letters":
		return namegen.NewThreeLetters(), "three_letters", nil
	case "cities":
		return namegen.NewCities(), "cities", nil
	case "countries":
		return namegen.NewCountries(), "countries", nil
	case "first_names":
		return namegen.NewFirstNames(), "first_names", nil
	default:
		return nil, "", fmt.Errorf("unknown name_generator_id: %s", generatorId)
	}
}
