package middleware

import (
	"Kubessistant/backend/endpoint"
	"Kubessistant/backend/objects"
)

type GeneralMiddleware struct {
	endpoint endpoint.IGeneralEndpoint
}

func NewGeneralMiddleware(endpoint endpoint.IGeneralEndpoint) *GeneralMiddleware {
	return &GeneralMiddleware{endpoint}
}

func (n *GeneralMiddleware) GetNodes() []objects.NodeDto {
	return n.endpoint.GetNodes()
}

func (n *GeneralMiddleware) GetNodeByName(name string) objects.NodeDto {
	return n.endpoint.GetNodeByName(name)
}
