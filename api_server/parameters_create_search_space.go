package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type CreateSearchSpaceParameters struct {
	Id                   string                           `json:"id"`
	Name                 string                           `json:"name"`
	GenerationDate       time.Time                        `json:"generation_date"`
	GenerationDurationMs int                              `json:"generation_duration_ms"`
	GenerationParameters GenerationJobParameters          `json:"generation_job_parameters"`
	Nodes                map[string]CreateSearchSpaceNode `json:"nodes"`
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
