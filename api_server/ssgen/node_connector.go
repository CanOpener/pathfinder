package ssgen

import (
	"fmt"

	"github.com/canopener/pathfinder/api_server/nodeconnect"
)

type nodeConnector struct {
	parameters GenerationJobParameters
}

type rawNodeConnector interface {
	ConnectNodes() (map[string]map[string]float64, error)
}

type fullNodeConnectorParameters struct {
	parameters GenerationJobParameters
	nodes      []nodeconnect.NodeDescriptor
}

func (p fullNodeConnectorParameters) GetMaximumNodeConnectionCount() int {
	return p.parameters.GetMaximumNodeConnectionCount()
}
func (p fullNodeConnectorParameters) GetNodes() []nodeconnect.NodeDescriptor {
	return p.nodes
}

func newNodeConnector(parameters GenerationJobParameters) (nodeConnector, error) {
	return nodeConnector{parameters: parameters}, nil
}

func (c nodeConnector) connectNodes(nodes map[string]Node) (map[string]Node, error) {
	rawNodeConnector, err := c.initRawNodeConnector(nodes)
	if err != nil {
		return nil, err
	}

	nodeConnections, err := rawNodeConnector.ConnectNodes()
	if err != nil {
		return nil, err
	}

	newNodes := map[string]Node{}
	for targetNodeId, targetNode := range nodes {
		newNode := Node{
			X:           targetNode.X,
			Y:           targetNode.Y,
			Connections: nodeConnections[targetNodeId],
		}
		newNodes[targetNodeId] = newNode
	}

	return newNodes, nil
}

func (c nodeConnector) initRawNodeConnector(nodes map[string]Node) (rawNodeConnector, error) {
	nodeDescriptors := []nodeconnect.NodeDescriptor{}
	for id, node := range nodes {
		nodeDescriptors = append(nodeDescriptors, nodeconnect.NodeDescriptor{
			X:  node.X,
			Y:  node.Y,
			Id: id,
		})
	}
	fullParameters := fullNodeConnectorParameters{
		parameters: c.parameters,
		nodes:      nodeDescriptors,
	}

	switch c.parameters.GetNodeConnectorId() {
	case "prim":
		return nodeconnect.NewPrim(fullParameters), nil
	case "maxn":
		return nodeconnect.NewMaxN(fullParameters), nil
	case "none":
		return nodeconnect.NewNone(fullParameters), nil
	default:
		return nil, fmt.Errorf("unknown node connector: %s", c.parameters.GetNodeConnectorId())
	}
}
