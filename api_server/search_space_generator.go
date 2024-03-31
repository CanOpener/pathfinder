package main

import (
	"fmt"

	"github.com/canopener/pathfinder/api_server/ssgen"
)

type searchSpaceGenerator interface {
	Generate(GenerationJobParameters) (searchSpaceDescriptor, error)
}

type searchSpaceDescriptor interface{}

type ssgenGenerator struct {
	generator ssgen.Generator
}

func newSearchSpaceGenerator(generatorId string) (searchSpaceGenerator, error) {
	switch generatorId {
	case "ssgen":
		return newSsgenGenerator()
	default:
		return nil, fmt.Errorf("invalid search generator id: %s", generatorId)
	}
}

func newSsgenGenerator() (*ssgenGenerator, error) {
	return &ssgenGenerator{
		generator: ssgen.Generator{},
	}, nil
}

func (g *ssgenGenerator) Generate(parameters GenerationJobParameters) (searchSpaceDescriptor, error) {
	generator, err := ssgen.NewGenerator(parameters)
	if err != nil {
		return nil, err
	}
	g.generator = generator
	searchSpace, err := g.generator.Generate()

	return searchSpace, err
}
