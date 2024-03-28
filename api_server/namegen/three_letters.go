package namegen

import (
	"fmt"
	"math/rand"
)

type ThreeLetters struct {
	alphabet [26]string
}

func NewThreeLetters() ThreeLetters {
	return ThreeLetters{
		alphabet: [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
			"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"},
	}
}

func (g ThreeLetters) NewName() string {
	randomIndex1 := rand.Intn(len(g.alphabet))
	randomIndex2 := rand.Intn(len(g.alphabet))
	randomIndex3 := rand.Intn(len(g.alphabet))
	return fmt.Sprintf("%s%s%s", g.alphabet[randomIndex1], g.alphabet[randomIndex2],
		g.alphabet[randomIndex3])
}
