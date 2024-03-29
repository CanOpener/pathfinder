package ssgen

import (
	"time"

	"github.com/google/uuid"
)

type SearchSpace struct {
	Id                   string               `json:"id"`
	GenerationDate       time.Time            `json:"generation_date"`
	GenerationParameters GenerationParameters `json:"generation_parameters"`
	Nodes                map[string]Node      `json:"nodes"`
}

func newSearchSpace(parameters GenerationParameters) SearchSpace {
	return SearchSpace{
		Id:                   uuid.New().String(),
		GenerationDate:       time.Now().UTC(),
		GenerationParameters: parameters,
		Nodes:                map[string]Node{},
	}
}
