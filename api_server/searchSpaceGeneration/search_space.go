package searchSpaceGeneration

import (
	"time"

	"github.com/google/uuid"
)

type SearchSpace struct {
	Id             string          `json:"id"`
	GenerationDate time.Time       `json:"generation_date"`
	Algorithm      string          `json:"algorithm"`
	NameGenerator  string          `json:"name_generator"`
	NodeCount      int             `json:"node_count"`
	Nodes          map[string]Node `json:"nodes"`
}

type GenerationParameters interface {
	Algorithm() string
	NameGenerator() string
	NodeCount() int
}

func newSearchSpace(params GenerationParameters) SearchSpace {
	return SearchSpace{
		Id:             uuid.New().String(),
		GenerationDate: time.Now(),
		Algorithm:      params.Algorithm(),
		NameGenerator:  params.NameGenerator(),
		NodeCount:      params.NodeCount(),
		Nodes:          map[string]Node{},
	}
}
