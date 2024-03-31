package ssgen

import (
	"fmt"
	"strconv"

	"github.com/canopener/pathfinder/api_server/nodeplot"
)

type nodePlotter struct {
	generator rawNodePlotter
}
type rawNodePlotter interface {
	PlotNodes() ([]nodeplot.Coordinate, error)
}

func newNodePlotter(parameters GenerationParameters) (nodePlotter, error) {
	switch parameters.GetNodePlotterId() {
	case "random":
		return nodePlotter{generator: nodeplot.NewRandom(parameters)}, nil
	default:
		return nodePlotter{}, fmt.Errorf("unknown node plotter: %s", parameters.GetNodePlotterId())
	}
}

func (p nodePlotter) plotNodes() (map[string]Node, error) {
	nodeCoordinates, err := p.generator.PlotNodes()
	if err != nil {
		return nil, err
	}

	nodeCount := 0
	newNodes := map[string]Node{}
	for _, nodeCoordinate := range nodeCoordinates {
		newNode := Node{
			X:           nodeCoordinate.X,
			Y:           nodeCoordinate.Y,
			Connections: map[string]float64{},
		}
		newNodes[strconv.Itoa(nodeCount)] = newNode
		nodeCount += 1
	}
	return newNodes, nil
}
