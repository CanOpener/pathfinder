package nodeconnect

import (
	"errors"
	"fmt"
)

type Prim struct {
	parameters   PrimParameters
	nodes        map[string]node
	visitedNodes map[string]bool
}

type connection struct {
	visitedNodeId string
	newNodeId     string
	distance      float64
}

type PrimParameters interface {
	GetNodes() []NodeDescriptor
}

func NewPrim(parameters PrimParameters) *Prim {
	return &Prim{
		parameters:   parameters,
		nodes:        nodeSet(parameters),
		visitedNodes: map[string]bool{},
	}
}

func (c *Prim) ConnectNodes() (map[string]map[string]float64, error) {
	_, randomNodeId, err := c.anyNode()
	if err != nil {
		return nil, fmt.Errorf("prim node connector failed due to: %w", err)
	}

	c.visitedNodes[randomNodeId] = true
	for len(c.visitedNodes) < len(c.nodes) {
		shortestConnection, err := c.shortestPossibleConnection()
		if err != nil {
			return nil, fmt.Errorf("prim node connector failed due to: %w", err)
		}

		visitedNode := c.nodes[shortestConnection.visitedNodeId]
		newNode := c.nodes[shortestConnection.newNodeId]
		connectionDistance := visitedNode.distanceTo(newNode)

		visitedNode.connections[shortestConnection.newNodeId] = connectionDistance
		c.nodes[shortestConnection.visitedNodeId] = visitedNode

		newNode.connections[shortestConnection.visitedNodeId] = connectionDistance
		c.nodes[shortestConnection.newNodeId] = newNode

		c.visitedNodes[shortestConnection.newNodeId] = true
	}

	return serializeConnections(c.nodes), nil
}

func (c *Prim) anyNode() (node, string, error) {
	for nodeId, node := range c.nodes {
		return node, nodeId, nil
	}
	return node{}, "", errors.New("empty set of all nodes")
}

func (c *Prim) shortestPossibleConnection() (connection, error) {
	shortestPossibleConnection := connection{}
	oneConnectionFound := false

	for visitedNodeId, _ := range c.visitedNodes {
		shortestNodeConnection, ok := c.shortestPossibleConnectionForNode(visitedNodeId)
		if !ok {
			continue
		}

		if (!oneConnectionFound) || (shortestNodeConnection.distance < shortestPossibleConnection.distance) {
			shortestPossibleConnection = shortestNodeConnection
			oneConnectionFound = true
		}
	}

	if oneConnectionFound {
		return shortestPossibleConnection, nil
	}
	return shortestPossibleConnection, errors.New("no possible connections")
}

func (c *Prim) shortestPossibleConnectionForNode(targetNodeId string) (connection, bool) {
	_, ok := c.visitedNodes[targetNodeId]
	if !ok {
		return connection{}, false
	}

	targetNode := c.nodes[targetNodeId]
	shortestPossibleConnection := connection{}
	oneConnectionFound := false

	for id, potentialConnectingNode := range c.nodes {
		_, ok := c.visitedNodes[id]
		if ok {
			continue
		}

		possibleConnection := connection{
			visitedNodeId: targetNodeId,
			newNodeId:     id,
			distance:      targetNode.distanceTo(potentialConnectingNode),
		}
		if (!oneConnectionFound) || (possibleConnection.distance < shortestPossibleConnection.distance) {
			shortestPossibleConnection = possibleConnection
			oneConnectionFound = true
		}
	}

	return shortestPossibleConnection, oneConnectionFound
}
