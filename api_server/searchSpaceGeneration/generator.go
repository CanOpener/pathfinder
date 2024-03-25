package searchSpaceGeneration

import (
	"fmt"
)

func (ss *SearchSpace) generate() error {
	err := ss.generateUnconnectedNodes()
	if err != nil {
		return fmt.Errorf("unconnected node generation error: %w", err)
	}

	switch ss.Algorithm {
	case "default":
		return ss.generatePrim()
	case "prim":
		return ss.generatePrim()
	default:
		return fmt.Errorf("unknown algorithm: %s", ss.Algorithm)
	}
}

func (ss *SearchSpace) generateUnconnectedNodes() error {
	if ss.NodeCount <= 0 || ss.NodeCount >= 10000 {
		return fmt.Errorf("node_count must be between 1 and 9999 (inclusive): %d", ss.NodeCount)
	}
	nodeNameGenerator, err := ss.nameGenerator()
	if err != nil {
		return err
	}

	for len(ss.Nodes) < ss.NodeCount {
		node := newNode(nodeNameGenerator.NewName())
		if ss.nodeDistanceValid(node) && ss.nodeNameValid(node) {
			ss.Nodes[node.Id] = node
		}
	}

	return nil
}

func (ss *SearchSpace) nodeDistanceValid(node Node) bool {
	for _, targetNode := range ss.Nodes {
		if node.distanceTo(targetNode) < 25 {
			return false
		}
	}

	return true
}

func (ss *SearchSpace) nodeNameValid(node Node) bool {
	_, ok := ss.Nodes[node.Id]
	return !ok
}

func Generate(params GenerationParameters) (SearchSpace, error) {
	searchSpace := newSearchSpace(params)
	err := searchSpace.generate()
	if err != nil {
		return searchSpace, err
	}

	return searchSpace, nil
}
