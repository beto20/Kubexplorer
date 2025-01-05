package middleware

import (
	"Kubessistant/backend/endpoint"
	"Kubessistant/backend/objects"
)

type NodeMiddleware struct {
	endpoint endpoint.INodeEndpoint
}

func NewNodeMiddleware(endpoint endpoint.INodeEndpoint) *NodeMiddleware {
	return &NodeMiddleware{endpoint}
}

func (n *NodeMiddleware) GetNodes() []objects.NodeDto {
	return n.endpoint.GetNodes()
}
