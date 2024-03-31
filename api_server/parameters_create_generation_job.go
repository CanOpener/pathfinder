package main

import (
	"encoding/json"
	"net/http"
)

type GenerationJobParameters struct {
	NodePlotterParameters   NodePlotterParameters   `json:"node_plotter_parameters"`
	NodeConnectorParameters NodeConnectorParameters `json:"node_connector_parameters"`
	NameGeneratorParameters NameGeneratorParameters `json:"name_generator_parameters"`
}

type NodePlotterParameters struct {
	NodePlotterId       string `json:"node_plotter_id"`
	NodeCount           int    `json:"node_count"`
	MinimumDistance     int    `json:"minimum_distance"`
	MaximumPlotAttempts int    `json:"maximum_plot_attempts"`
	GridSizeX           int    `json:"grid_size_x"`
	GridSizeY           int    `json:"grid_size_y"`
}

type NodeConnectorParameters struct {
	NodeConnectorId            string `json:"node_connector_id"`
	MaximumNodeConnectionCount int    `json:"maximum_node_connection_count"`
}

type NameGeneratorParameters struct {
	NameGeneratorId       string `json:"name_generator_id"`
	AllowDuplicates       bool   `json:"allow_duplicates"`
	MaximumSampleAttempts int    `json:"maximum_sample_attempts"`
}

func (p GenerationJobParameters) GetNodePlotterId() string {
	return p.NodePlotterParameters.NodePlotterId
}
func (p GenerationJobParameters) GetNodeCount() int {
	return p.NodePlotterParameters.NodeCount
}
func (p GenerationJobParameters) GetMinimumDistance() int {
	return p.NodePlotterParameters.MinimumDistance
}
func (p GenerationJobParameters) GetMaximumPlotAttempts() int {
	return p.NodePlotterParameters.MaximumPlotAttempts
}
func (p GenerationJobParameters) GetGridSizeX() int {
	return p.NodePlotterParameters.GridSizeX
}
func (p GenerationJobParameters) GetGridSizeY() int {
	return p.NodePlotterParameters.GridSizeY
}
func (p GenerationJobParameters) GetNodeConnectorId() string {
	return p.NodeConnectorParameters.NodeConnectorId
}
func (p GenerationJobParameters) GetMaximumNodeConnectionCount() int {
	return p.NodeConnectorParameters.MaximumNodeConnectionCount
}
func (p GenerationJobParameters) GetNameGeneratorId() string {
	return p.NameGeneratorParameters.NameGeneratorId
}
func (p GenerationJobParameters) GetAllowDuplicates() bool {
	return p.NameGeneratorParameters.AllowDuplicates
}
func (p GenerationJobParameters) GetMaximumSampleAttempts() int {
	return p.NameGeneratorParameters.MaximumSampleAttempts
}

func generationJobParametersFromRequest(request *http.Request) (GenerationJobParameters, error) {
	parameters := GenerationJobParameters{}
	err := json.NewDecoder(request.Body).Decode(&parameters)
	if err != nil {
		return parameters, err
	}

	return parameters, nil
}
