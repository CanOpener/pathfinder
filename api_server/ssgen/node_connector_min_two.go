package ssgen

import (
	"errors"
	"fmt"
	"math/rand"
)

type NodeConnectorMinTwo struct {
	nodes map[string]Node
}

func newNodeConnectorMinTwo() *NodeConnectorMinTwo {
	return &NodeConnectorMinTwo{
		nodes: map[string]Node{},
	}
}

func (c *NodeConnectorMinTwo) ConnectedNodes(nodes map[string]Node) (map[string]Node, error) {
	c.copyNodes(nodes)

	for id, node := range c.nodes {
		for len(node.Connections) < 2 {
			connection, err := c.randomNode(nodes)
			if err != nil {
				return nodes, err
			}

			if connection.Id == id {
				continue
			}

			node.Connections = append(node.Connections, connection.Id)
			c.nodes[id] = node
			connection.Connections = append(connection.Connections, id)
			c.nodes[connection.Id] = connection
		}
	}

	return c.nodes, nil
}

func (c *NodeConnectorMinTwo) copyNodes(nodes map[string]Node) {
	for id, node := range nodes {
		c.nodes[id] = node
	}
}

func (c *NodeConnectorMinTwo) randomNode(nodes map[string]Node) (Node, error) {
	if len(nodes) == 0 {
		return Node{}, errors.New("no nodes available")
	}

	index := rand.Intn(len(nodes))
	i := 0
	for _, node := range nodes {
		if i == index {
			return node, nil
		}
		i++
	}
	return Node{}, fmt.Errorf("failed to get random node at index: %d", index)
}
