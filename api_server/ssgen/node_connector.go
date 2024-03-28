package ssgen

import "fmt"

type nodeConnector interface {
	ConnectedNodes(map[string]Node) (map[string]Node, error)
}

func newNodeConnector(algorithmId string) (nodeConnector, string, error) {
	switch algorithmId {
	case "default":
		return newNodeConnectorPrim(), "prim", nil
	case "prim":
		return newNodeConnectorPrim(), "prim", nil
	case "min_two_conn":
		return newNodeConnectorMinTwo(), "min_two_conn", nil
	case "none":
		return newNodeConnectorNone(), "none", nil
	default:
		return nil, "", fmt.Errorf("unknown algorithm: %s", algorithmId)
	}
}
