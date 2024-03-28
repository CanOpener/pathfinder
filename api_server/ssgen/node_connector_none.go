package ssgen

type NodeConnectorNone struct{}

func newNodeConnectorNone() *NodeConnectorNone {
	return &NodeConnectorNone{}
}

func (c *NodeConnectorNone) ConnectedNodes(nodes map[string]Node) (map[string]Node, error) {
	return nodes, nil
}
