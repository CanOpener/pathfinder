package searchSpaceGeneration

import (
	"fmt"
	"math/rand"
)

type threeLetterNameGenerator struct {
	names []string
}

func newThreeLetterNameGenerator() threeLetterNameGenerator {
	return threeLetterNameGenerator{
		names: []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
			"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"},
	}
}

func (ng threeLetterNameGenerator) NewName() string {
	randomIndex1 := rand.Intn(len(ng.names))
	randomIndex2 := rand.Intn(len(ng.names))
	randomIndex3 := rand.Intn(len(ng.names))
	return fmt.Sprintf("%s%s%s", ng.names[randomIndex1], ng.names[randomIndex2], ng.names[randomIndex3])
}
