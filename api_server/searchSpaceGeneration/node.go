package searchSpaceGeneration

import (
	"math"
	"math/rand"
)

type Node struct {
	Id          string   `json:"id"`
	X           int      `json:"x"`
	Y           int      `json:"y"`
	Connections []string `json:"connections"`
}

func newNode(id string) Node {
	return Node{
		Id:          id,
		X:           rand.Intn(1000),
		Y:           rand.Intn(1000),
		Connections: []string{},
	}
}

func (n Node) distanceTo(otherNode Node) float64 {
	dx := float64(n.X - otherNode.X)
	dy := float64(n.Y - otherNode.Y)
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}
