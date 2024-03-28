package ssgen

import (
	"errors"
	"fmt"
)

type NodeConnectorPrim struct {
	nodes        map[string]Node
	visitedNodes map[string]bool
}

type connection struct {
	visitedNodeId string
	newNodeId     string
	distance      float64
}

func newNodeConnectorPrim() *NodeConnectorPrim {
	return &NodeConnectorPrim{
		nodes:        map[string]Node{},
		visitedNodes: map[string]bool{},
	}
}

func (c *NodeConnectorPrim) ConnectedNodes(nodes map[string]Node) (map[string]Node, error) {
	c.copyNodes(nodes)

	randomNode, err := c.anyNode()
	if err != nil {
		return nil, fmt.Errorf("prim node connector failed due to: %w", err)
	}

	c.visitedNodes[randomNode.Id] = true
	for len(c.visitedNodes) < len(c.nodes) {
		shortestConnection, err := c.shortestPossibleConnection()
		if err != nil {
			return nil, fmt.Errorf("prim node connector failed due to: %w", err)
		}

		visitedNode := c.nodes[shortestConnection.visitedNodeId]
		visitedNode.Connections = append(visitedNode.Connections, shortestConnection.newNodeId)
		c.nodes[visitedNode.Id] = visitedNode

		newNode := c.nodes[shortestConnection.newNodeId]
		newNode.Connections = append(newNode.Connections, shortestConnection.visitedNodeId)
		c.nodes[newNode.Id] = newNode

		c.visitedNodes[newNode.Id] = true
	}

	return c.nodes, nil
}

func (c *NodeConnectorPrim) copyNodes(nodes map[string]Node) {
	for id, node := range nodes {
		c.nodes[id] = node
	}
}

func (c *NodeConnectorPrim) anyNode() (Node, error) {
	for _, node := range c.nodes {
		return node, nil
	}
	return Node{}, errors.New("empty set of all nodes")
}

func (c *NodeConnectorPrim) shortestPossibleConnection() (connection, error) {
	shortestPossibleConnection := connection{}
	oneConnectionFound := false

	for visitedNodeId, _ := range c.visitedNodes {
		visitedNode, ok := c.nodes[visitedNodeId]
		if !ok {
			return connection{}, fmt.Errorf("visited node id not in set of all nodes: %s", visitedNodeId)
		}

		shortestNodeConnection, ok := c.shortestPossibleConnectionForNode(visitedNode)
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

func (c *NodeConnectorPrim) shortestPossibleConnectionForNode(node Node) (connection, bool) {
	_, ok := c.visitedNodes[node.Id]
	if !ok {
		return connection{}, false
	}

	shortestPossibleConnection := connection{}
	oneConnectionFound := false

	for id, potentialConnectingNode := range c.nodes {
		_, ok := c.visitedNodes[id]
		if ok {
			continue
		}

		possibleConnection := connection{
			visitedNodeId: node.Id,
			newNodeId:     id,
			distance:      node.distanceTo(potentialConnectingNode),
		}
		if (!oneConnectionFound) || (possibleConnection.distance < shortestPossibleConnection.distance) {
			shortestPossibleConnection = possibleConnection
			oneConnectionFound = true
		}
	}

	return shortestPossibleConnection, oneConnectionFound
}
