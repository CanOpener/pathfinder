package nodeplot

import (
	"fmt"
	"math/rand"
)

type Random struct {
	parameters RandomParameters
	nodes      []Coordinate
}

type RandomParameters interface {
	GetNodeCount() int
	GetMinimumDistance() int
	GetMaximumPlotAttempts() int
	GetGridSizeX() int
	GetGridSizeY() int
}

func NewRandom(parameters RandomParameters) *Random {
	return &Random{
		parameters: parameters,
		nodes:      []Coordinate{},
	}
}

func (p *Random) PlotNodes() ([]Coordinate, error) {
	for len(p.nodes) < p.parameters.GetNodeCount() {
		node, err := p.generateNode()
		if err != nil {
			return nil, err
		}

		p.nodes = append(p.nodes, node)
	}

	return p.nodes, nil
}

func (p *Random) generateNode() (Coordinate, error) {
	plotAttempts := 0
	for plotAttempts < p.parameters.GetMaximumPlotAttempts() {
		node := Coordinate{
			X: rand.Intn(p.parameters.GetGridSizeX()),
			Y: rand.Intn(p.parameters.GetGridSizeY()),
		}
		if p.nodeDistanceValid(node) {
			return node, nil
		}

		plotAttempts += 1
	}

	return Coordinate{}, fmt.Errorf("failed to generate random node with minimum distance %d to other nodes. attempt count :%d",
		p.parameters.GetMinimumDistance(), p.parameters.GetMaximumPlotAttempts())
}

func (p *Random) nodeDistanceValid(node Coordinate) bool {
	for _, targetNode := range p.nodes {
		if node.distanceTo(targetNode) < float64(p.parameters.GetMinimumDistance()) {
			return false
		}
	}

	return true
}
