package ssgen

import (
	"time"

	"github.com/google/uuid"
)

type SearchSpace struct {
	Name                    string                  `json:"name"`
	GenerationDate          time.Time               `json:"generation_date"`
	GenerationDurationMs    int                     `json:"generation_duration_ms"`
	GenerationJobParameters GenerationJobParameters `json:"generation_job_parameters"`
	GridSizeX               int                     `json:"grid_size_x"`
	GridSizeY               int                     `json:"grid_size_y"`
	Nodes                   map[string]Node         `json:"nodes"`
}

func newSearchSpace(parameters GenerationJobParameters) SearchSpace {
	return SearchSpace{
		Name:                    uuid.New().String(),
		GenerationDate:          time.Now().UTC(),
		GenerationDurationMs:    0,
		GridSizeX:               parameters.GetGridSizeX(),
		GridSizeY:               parameters.GetGridSizeY(),
		GenerationJobParameters: parameters,
		Nodes:                   map[string]Node{},
	}
}
