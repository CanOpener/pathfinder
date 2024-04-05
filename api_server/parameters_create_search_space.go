package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type CreateSearchSpaceParameters struct {
	Name                    string                           `json:"name"`
	GenerationDate          time.Time                        `json:"generation_date"`
	GenerationDurationMs    int                              `json:"generation_duration_ms"`
	GenerationJobParameters GenerationJobParameters          `json:"generation_job_parameters"`
	GridSizeX               int                              `json:"grid_size_x"`
	GridSizeY               int                              `json:"grid_size_y"`
	Nodes                   map[string]CreateSearchSpaceNode `json:"nodes"`
}

type CreateSearchSpaceNode struct {
	X           int                `json:"x"`
	Y           int                `json:"y"`
	Connections map[string]float64 `json:"connections"`
}

func createSearchSpaceParametersFromRequest(request *http.Request) (CreateSearchSpaceParameters, error) {
	parameters := CreateSearchSpaceParameters{}
	err := json.NewDecoder(request.Body).Decode(&parameters)
	if err != nil {
		return parameters, err
	}

	return parameters, nil
}
