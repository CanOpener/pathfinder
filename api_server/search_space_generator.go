package main

import (
	"fmt"
)

type searchSpaceGenerator interface {
	Generate() (searchSpaceDescriptor, int, error)
}
type searchSpaceDescriptor interface{}
type searchSpaceGenerationParameters struct {
	NodeConnectorId     string `json:"node_connector_id"`
	NameGeneratorId     string `json:"name_generator_id"`
	NodeCount           int    `json:"node_count"`
	MinimumNodeDistance int    `json:"minimum_node_distance"`
}

func (p searchSpaceGenerationParameters) GetNameGeneratorId() string  { return p.NameGeneratorId }
func (p searchSpaceGenerationParameters) GetNodeConnectorId() string  { return p.NodeConnectorId }
func (p searchSpaceGenerationParameters) GetNodeCount() int           { return p.NodeCount }
func (p searchSpaceGenerationParameters) GetMinimumNodeDistance() int { return p.MinimumNodeDistance }

func newSearchSpaceGenerator(generator_id string, parameters searchSpaceGenerationParameters) (searchSpaceGenerator, error) {
	switch generator_id {
	case "ssgen":
		return newSsgenGenerator(parameters)
	default:
		return nil, fmt.Errorf("invalid search generator id: %s", generator_id)
	}
}
