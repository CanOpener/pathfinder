package main

import (
	"time"

	"github.com/canopener/pathfinder/api_server/ssgen"
)

type ssgenGenerator struct {
	generator ssgen.Generator
}

func newSsgenGenerator() (*ssgenGenerator, error) {
	return &ssgenGenerator{
		generator: ssgen.Generator{},
	}, nil
}

func (g *ssgenGenerator) Generate(parameters SearchSpaceGenerationParameters) (searchSpaceDescriptor, int, error) {
	generator, err := ssgen.NewGenerator(parameters)
	if err != nil {
		return nil, 0, err
	}
	g.generator = generator

	startTime := time.Now()
	searchSpace, err := g.generator.Generate()
	endTime := time.Now()

	return searchSpace, int(endTime.Sub(startTime).Milliseconds()), err
}
