package ssgen

import (
	"fmt"
)

type Generator struct {
	parameters    GenerationParameters
	nodeConnector nodeConnector
	nameGenerator nameGenerator
	searchSpace   SearchSpace
}

type GenerationParameters interface {
	GetNodeConnectorId() string
	GetNameGeneratorId() string
	GetNodeCount() int
	GetMinimumNodeDistance() int
}

func NewGenerator(parameters GenerationParameters) (Generator, error) {
	nodeConnector, err := newNodeConnector(parameters.GetNodeConnectorId())
	if err != nil {
		return Generator{}, err
	}

	nameGenerator, err := newNameGenerator(parameters.GetNameGeneratorId())
	if err != nil {
		return Generator{}, err
	}

	return Generator{
		parameters:    parameters,
		nodeConnector: nodeConnector,
		nameGenerator: nameGenerator,
		searchSpace:   newSearchSpace(parameters),
	}, nil
}

func (g *Generator) Generate() (SearchSpace, error) {
	err := g.plotNodes()
	if err != nil {
		return SearchSpace{}, err
	}

	err = g.connectNodes()
	if err != nil {
		return SearchSpace{}, err
	}

	return g.searchSpace, nil
}

func (g *Generator) plotNodes() error {
	if g.parameters.GetNodeCount() <= 0 || g.parameters.GetNodeCount() >= 10000 {
		return fmt.Errorf("node_count must be between 1 and 9999 (inclusive): %d", g.parameters.GetNodeCount())
	}

	for len(g.searchSpace.Nodes) < g.parameters.GetNodeCount() {
		node := newNode(g.nameGenerator.NewName())
		if g.nodeDistanceValid(node) && g.nodeNameValid(node) {
			g.searchSpace.Nodes[node.Id] = node
		}
	}

	return nil
}

func (g *Generator) nodeDistanceValid(node Node) bool {
	for _, targetNode := range g.searchSpace.Nodes {
		if node.distanceTo(targetNode) < float64(g.parameters.GetMinimumNodeDistance()) {
			return false
		}
	}

	return true
}

func (g *Generator) nodeNameValid(node Node) bool {
	_, ok := g.searchSpace.Nodes[node.Id]
	return !ok
}

func (g *Generator) connectNodes() error {
	connectedNodes, err := g.nodeConnector.ConnectedNodes(g.searchSpace.Nodes)
	if err != nil {
		return err
	}
	g.searchSpace.Nodes = connectedNodes
	return nil
}
