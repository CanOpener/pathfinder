package nodeconnect

import "sort"

type MaxN struct {
	parameters MaxNParameters
	nodes      map[string]node
}

type MaxNParameters interface {
	GetMaximumNodeConnectionCount() int
	GetNodes() []NodeDescriptor
}

func NewMaxN(parameters MaxNParameters) *MaxN {
	return &MaxN{
		parameters: parameters,
		nodes:      nodeSet(parameters),
	}
}

func (c *MaxN) ConnectNodes() (map[string]map[string]float64, error) {
	for originalNodeId, originalNode := range c.nodes {
		for _, connectingNodeId := range c.closestNodeIds(originalNodeId) {
			if len(originalNode.connections) >= c.parameters.GetMaximumNodeConnectionCount() {
				break
			}

			if originalNodeId == connectingNodeId {
				continue
			}

			if _, ok := originalNode.connections[connectingNodeId]; !ok {
				continue
			}

			connectingNode := c.nodes[connectingNodeId]
			if len(connectingNode.connections) >= c.parameters.GetMaximumNodeConnectionCount() {
				continue
			}

			connectionDistance := originalNode.distanceTo(connectingNode)
			originalNode.connections[connectingNodeId] = connectionDistance
			c.nodes[originalNodeId] = originalNode
			connectingNode.connections[originalNodeId] = connectionDistance
			c.nodes[connectingNodeId] = connectingNode
		}
	}

	return serializeConnections(c.nodes), nil
}

func (c *MaxN) closestNodeIds(nodeId string) []string {
	allOtherNodeIds := c.otherNodeIds(nodeId)
	targetNode := c.nodes[nodeId]
	sort.Slice(allOtherNodeIds, func(i, j int) bool {
		return c.nodes[allOtherNodeIds[i]].distanceTo(targetNode) <
			c.nodes[allOtherNodeIds[j]].distanceTo(targetNode)
	})
	return allOtherNodeIds
}

func (c *MaxN) otherNodeIds(nodeId string) []string {
	ids := []string{}
	for otherNodeId := range c.nodes {
		if nodeId == otherNodeId {
			continue
		}

		ids = append(ids, otherNodeId)
	}
	return ids
}
