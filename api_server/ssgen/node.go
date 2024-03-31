package ssgen

import (
	"math"
)

type Node struct {
	X           int                `json:"x"`
	Y           int                `json:"y"`
	Connections map[string]float64 `json:"connections"`
}

func (n Node) distanceTo(otherNode Node) float64 {
	dx := float64(n.X - otherNode.X)
	dy := float64(n.Y - otherNode.Y)
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}
