package nodeconnect

import "math"

type node struct {
	x           int
	y           int
	connections map[string]float64
}

type parametersWithNodes interface {
	GetNodes() []NodeDescriptor
}

type NodeDescriptor struct {
	Id string
	X  int
	Y  int
}

func (n node) distanceTo(node node) float64 {
	dx := float64(n.x - node.x)
	dy := float64(n.y - node.y)
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}

func serializeConnections(nodes map[string]node) map[string]map[string]float64 {
	connections := map[string]map[string]float64{}
	for targetNodeId, targetNode := range nodes {
		connections[targetNodeId] = map[string]float64{}
		for connectingNodeId := range targetNode.connections {
			connections[targetNodeId][connectingNodeId] = targetNode.connections[connectingNodeId]
		}
	}
	return connections
}

func nodeSet(parameters parametersWithNodes) map[string]node {
	nodes := map[string]node{}
	for _, targetNode := range parameters.GetNodes() {
		nodes[targetNode.Id] = node{
			x:           targetNode.X,
			y:           targetNode.Y,
			connections: map[string]float64{},
		}
	}
	return nodes
}
