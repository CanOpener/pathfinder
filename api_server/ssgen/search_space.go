package ssgen

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type SearchSpace struct {
	Id              string               `json:"id"`
	Parameters      GenerationParameters `json:"parameters"`
	GenerationDate  time.Time            `json:"generation_date"`
	nodeConnector   nodeConnector        `json:"-"`
	NodeConnectorId string               `json:"node_connector_id"`
	nameGenerator   nameGenerator        `json:"-"`
	NameGeneratorId string               `json:"name_generator_id"`
	Nodes           map[string]Node      `json:"nodes"`
}

type GenerationParameters interface {
	GetNodeConnectorId() string
	GetNameGeneratorId() string
	GetNodeCount() int
	GetMinimumNodeDistance() int
}

func Generate(params GenerationParameters) (SearchSpace, error) {
	searchSpace, err := newSearchSpace(params)
	if err != nil {
		return SearchSpace{}, err
	}

	err = searchSpace.generateUnconnectedNodes()
	if err != nil {
		return SearchSpace{}, err
	}

	err = searchSpace.connectNodes()
	if err != nil {
		return SearchSpace{}, err
	}

	return searchSpace, nil
}

func newSearchSpace(params GenerationParameters) (SearchSpace, error) {
	nodeConnector, nodeConnectorId, err := newNodeConnector(params.GetNodeConnectorId())
	if err != nil {
		return SearchSpace{}, err
	}

	nameGenerator, nameGeneratorId, err := newNameGenerator(params.GetNameGeneratorId())
	if err != nil {
		return SearchSpace{}, err
	}

	return SearchSpace{
		Id:              uuid.New().String(),
		GenerationDate:  time.Now(),
		Parameters:      params,
		nodeConnector:   nodeConnector,
		NodeConnectorId: nodeConnectorId,
		nameGenerator:   nameGenerator,
		NameGeneratorId: nameGeneratorId,
		Nodes:           map[string]Node{},
	}, nil
}

func (ss *SearchSpace) generateUnconnectedNodes() error {
	if ss.Parameters.GetNodeCount() <= 0 || ss.Parameters.GetNodeCount() >= 10000 {
		return fmt.Errorf("node_count must be between 1 and 9999 (inclusive): %d", ss.Parameters.GetNodeCount())
	}

	for len(ss.Nodes) < ss.Parameters.GetNodeCount() {
		node := newNode(ss.nameGenerator.NewName())
		if ss.nodeDistanceValid(node) && ss.nodeNameValid(node) {
			ss.Nodes[node.Id] = node
		}
	}

	return nil
}

func (ss *SearchSpace) nodeDistanceValid(node Node) bool {
	for _, targetNode := range ss.Nodes {
		if node.distanceTo(targetNode) < float64(ss.Parameters.GetMinimumNodeDistance()) {
			return false
		}
	}

	return true
}

func (ss *SearchSpace) nodeNameValid(node Node) bool {
	_, ok := ss.Nodes[node.Id]
	return !ok
}

func (ss *SearchSpace) connectNodes() error {
	connectedNodes, err := ss.nodeConnector.ConnectedNodes(ss.Nodes)
	if err != nil {
		return err
	}
	ss.Nodes = connectedNodes
	return nil
}
