package ssgen

type Node struct {
	X           int                `json:"x"`
	Y           int                `json:"y"`
	Connections map[string]float64 `json:"connections"`
}
