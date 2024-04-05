package ssgen

import (
	"fmt"
	"time"
)

type Generator struct {
	parameters    GenerationJobParameters
	nodePlotter   nodePlotter
	nodeConnector nodeConnector
	nameGenerator nameGenerator
	searchSpace   SearchSpace
}

type GenerationJobParameters interface {
	GetNodePlotterId() string
	GetNodeCount() int
	GetMinimumDistance() int
	GetMaximumPlotAttempts() int
	GetGridSizeX() int
	GetGridSizeY() int
	GetNodeConnectorId() string
	GetMaximumNodeConnectionCount() int
	GetNameGeneratorId() string
	GetAllowDuplicates() bool
	GetMaximumSampleAttempts() int
}

func NewGenerator(parameters GenerationJobParameters) (Generator, error) {
	nodePlotter, err := newNodePlotter(parameters)
	if err != nil {
		return Generator{}, fmt.Errorf("failed to inint nodePlotter: %w", err)
	}

	nodeConnector, err := newNodeConnector(parameters)
	if err != nil {
		return Generator{}, fmt.Errorf("failed to inint nodeConnector: %w", err)
	}

	nameGenerator, err := newNameGenerator(parameters)
	if err != nil {
		return Generator{}, fmt.Errorf("failed to inint nameGenerator: %w", err)
	}

	return Generator{
		parameters:    parameters,
		nodePlotter:   nodePlotter,
		nodeConnector: nodeConnector,
		nameGenerator: nameGenerator,
		searchSpace:   newSearchSpace(parameters),
	}, nil
}

func (g *Generator) Generate() (SearchSpace, error) {
	startTime := time.Now()
	err := g.plotNodes()
	if err != nil {
		return SearchSpace{}, fmt.Errorf("ssgen failed to plot nodes :%w", err)
	}

	err = g.connectNodes()
	if err != nil {
		return SearchSpace{}, fmt.Errorf("ssgen failed to connect nodes :%w", err)
	}

	err = g.nameNodes()
	if err != nil {
		return SearchSpace{}, fmt.Errorf("ssgen failed to name nodes :%w", err)
	}

	g.searchSpace.GenerationDurationMs = int(time.Since(startTime).Milliseconds())
	return g.searchSpace, nil
}

func (g *Generator) plotNodes() error {
	nodes, err := g.nodePlotter.plotNodes()
	if err != nil {
		return err
	}
	g.searchSpace.Nodes = nodes
	return nil
}

func (g *Generator) connectNodes() error {
	connectedNodes, err := g.nodeConnector.connectNodes(g.searchSpace.Nodes)
	if err != nil {
		return err
	}
	g.searchSpace.Nodes = connectedNodes
	return nil
}

func (g *Generator) nameNodes() error {
	namedNodes, err := g.nameGenerator.nameNodes(g.searchSpace.Nodes)
	if err != nil {
		return err
	}
	g.searchSpace.Nodes = namedNodes
	return nil
}
