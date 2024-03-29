package ssgen

import "fmt"

type nodeConnector interface {
	ConnectedNodes(map[string]Node) (map[string]Node, error)
}

func newNodeConnector(algorithmId string) (nodeConnector, error) {
	switch algorithmId {
	case "prim":
		return newNodeConnectorPrim(), nil
	case "min_two_conn":
		return newNodeConnectorMinTwo(), nil
	case "none":
		return newNodeConnectorNone(), nil
	default:
		return nil, fmt.Errorf("unknown algorithm: %s", algorithmId)
	}
}
