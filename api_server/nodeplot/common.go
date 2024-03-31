package nodeplot

import "math"

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) distanceTo(coordinate Coordinate) float64 {
	dx := float64(c.X - coordinate.X)
	dy := float64(c.Y - coordinate.Y)
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}
