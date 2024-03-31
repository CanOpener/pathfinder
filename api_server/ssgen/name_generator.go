package ssgen

import (
	"fmt"

	"github.com/canopener/pathfinder/api_server/namegen"
)

type nameGenerator struct {
	generator rawNameGenerator
}
type rawNameGenerator interface {
	NewName() (string, error)
}

func newNameGenerator(parameters GenerationParameters) (nameGenerator, error) {
	switch parameters.GetNameGeneratorId() {
	case "three_letters":
		return nameGenerator{generator: namegen.NewThreeLetters(parameters)}, nil
	case "cities":
		return nameGenerator{generator: namegen.NewCities(parameters)}, nil
	case "countries":
		return nameGenerator{generator: namegen.NewCountries(parameters)}, nil
	case "first_names":
		return nameGenerator{generator: namegen.NewFirstNames(parameters)}, nil
	case "uuid":
		return nameGenerator{generator: namegen.NewUuid(parameters)}, nil
	default:
		return nameGenerator{}, fmt.Errorf("invalid name generator id: %s", parameters.GetNameGeneratorId())
	}
}

func (g nameGenerator) nameNodes(nodes map[string]Node) (map[string]Node, error) {
	oldToNewNameMap, err := g.mapNewNames(nodes)
	if err != nil {
		return nil, err
	}

	newNodes := map[string]Node{}
	for oldId, node := range nodes {
		newName := oldToNewNameMap[oldId]
		newNode := Node{
			X:           node.X,
			Y:           node.Y,
			Connections: map[string]float64{},
		}
		for oldConnectedNodeId, connectionDistance := range node.Connections {
			newConnectedNodeName := oldToNewNameMap[oldConnectedNodeId]
			newNode.Connections[newConnectedNodeName] = connectionDistance
		}
		newNodes[newName] = newNode
	}

	return newNodes, nil
}

func (g nameGenerator) mapNewNames(nodes map[string]Node) (map[string]string, error) {
	oldToNewNameMap := map[string]string{}
	for id, _ := range nodes {
		newName, err := g.generator.NewName()
		if err != nil {
			return nil, err
		}
		oldToNewNameMap[id] = newName
	}
	return oldToNewNameMap, nil
}
