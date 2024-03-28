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

func (a ThreeLetters) NewName() string {
	randomIndex1 := rand.Intn(len(a.alphabet))
	randomIndex2 := rand.Intn(len(a.alphabet))
	randomIndex3 := rand.Intn(len(a.alphabet))
	return fmt.Sprintf("%s%s%s", a.alphabet[randomIndex1], a.alphabet[randomIndex2],
		a.alphabet[randomIndex3])
}
