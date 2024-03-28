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
	default:
		return nil, "", fmt.Errorf("unknown name_generator_id: %s", generatorId)
	}
}
