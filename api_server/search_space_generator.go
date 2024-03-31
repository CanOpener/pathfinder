package main

import (
	"fmt"
)

type searchSpaceGenerator interface {
	Generate(SearchSpaceGenerationParameters) (searchSpaceDescriptor, int, error)
}

type searchSpaceDescriptor interface{}

type SearchSpaceGenerationParameters struct {
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

func (p SearchSpaceGenerationParameters) GetNodePlotterId() string {
	return p.NodePlotterParameters.NodePlotterId
}
func (p SearchSpaceGenerationParameters) GetNodeCount() int {
	return p.NodePlotterParameters.NodeCount
}
func (p SearchSpaceGenerationParameters) GetMinimumDistance() int {
	return p.NodePlotterParameters.MinimumDistance
}
func (p SearchSpaceGenerationParameters) GetMaximumPlotAttempts() int {
	return p.NodePlotterParameters.MaximumPlotAttempts
}
func (p SearchSpaceGenerationParameters) GetGridSizeX() int {
	return p.NodePlotterParameters.GridSizeX
}
func (p SearchSpaceGenerationParameters) GetGridSizeY() int {
	return p.NodePlotterParameters.GridSizeY
}
func (p SearchSpaceGenerationParameters) GetNodeConnectorId() string {
	return p.NodeConnectorParameters.NodeConnectorId
}
func (p SearchSpaceGenerationParameters) GetMaximumNodeConnectionCount() int {
	return p.NodeConnectorParameters.MaximumNodeConnectionCount
}
func (p SearchSpaceGenerationParameters) GetNameGeneratorId() string {
	return p.NameGeneratorParameters.NameGeneratorId
}
func (p SearchSpaceGenerationParameters) GetAllowDuplicates() bool {
	return p.NameGeneratorParameters.AllowDuplicates
}
func (p SearchSpaceGenerationParameters) GetMaximumSampleAttempts() int {
	return p.NameGeneratorParameters.MaximumSampleAttempts
}

func newSearchSpaceGenerator(generator_id string) (searchSpaceGenerator, error) {
	switch generator_id {
	case "ssgen":
		return newSsgenGenerator()
	default:
		return nil, fmt.Errorf("invalid search generator id: %s", generator_id)
	}
}
