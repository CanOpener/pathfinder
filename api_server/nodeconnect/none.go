package nodeconnect

type None struct {
	parameters NoneParameters
	nodes      map[string]node
}

type NoneParameters interface{}

func NewNone(parameters NoneParameters) *None {
	return &None{}
}

func (c *None) ConnectNodes() (map[string]map[string]float64, error) {
	return serializeConnections(c.nodes), nil
}
