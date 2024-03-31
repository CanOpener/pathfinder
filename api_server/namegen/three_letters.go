package namegen

import (
	"fmt"
	"math/rand"
)

type ThreeLetters struct {
	parameters NameGeneratorParameters
	takenNames map[string]int
	alphabet   [26]string
}

func NewThreeLetters(parameters NameGeneratorParameters) *ThreeLetters {
	return &ThreeLetters{
		parameters: parameters,
		takenNames: map[string]int{},
		alphabet: [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
			"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"},
	}
}

func (g *ThreeLetters) NewName() (string, error) {
	if g.parameters.GetAllowDuplicates() {
		return g.newNameWithDuplicates(), nil
	}
	return g.newNameWithoutDuplicates()
}

func (g *ThreeLetters) newNameWithDuplicates() string {
	originalNewName := g.randomName()
	finalNewName := originalNewName
	count := g.takenNames[originalNewName]
	if count > 0 {
		finalNewName = fmt.Sprintf("%s%d", originalNewName, count)
	}
	g.takenNames[originalNewName] += 1
	return finalNewName
}

func (g *ThreeLetters) newNameWithoutDuplicates() (string, error) {
	attemptCount := 0
	for attemptCount < g.parameters.GetMaximumSampleAttempts() {
		potentialName := g.randomName()
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

func (g *ThreeLetters) randomName() string {
	randomIndex1 := rand.Intn(len(g.alphabet))
	randomIndex2 := rand.Intn(len(g.alphabet))
	randomIndex3 := rand.Intn(len(g.alphabet))
	return fmt.Sprintf("%s%s%s", g.alphabet[randomIndex1], g.alphabet[randomIndex2],
		g.alphabet[randomIndex3])
}
