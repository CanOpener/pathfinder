package main

import (
	"time"

	"github.com/canopener/pathfinder/api_server/ssgen"
)

type ssgenGenerator struct {
	generator ssgen.Generator
}

func newSsgenGenerator(parameters searchSpaceGenerationParameters) (*ssgenGenerator, error) {
	generator, err := ssgen.NewGenerator(parameters)
	if err != nil {
		return nil, nil
	}

	return &ssgenGenerator{
		generator: generator,
	}, nil
}

func (g *ssgenGenerator) Generate() (searchSpaceDescriptor, int, error) {
	startTime := time.Now()
	searchSpace, err := g.generator.Generate()
	endTime := time.Now()

	return searchSpace, int(endTime.Sub(startTime).Milliseconds()), err
}
