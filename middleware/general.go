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

func (g *GeneralMiddleware) GetNodes() []objects.NodeDto {
	return g.endpoint.GetNodes()
}

func (g *GeneralMiddleware) GetNodeByName(name string) objects.NodeDto {
	return g.endpoint.GetNodeByName(name)
}

func (g *GeneralMiddleware) GetNamespace() []objects.NamespaceDto {
	return g.endpoint.GetNamespace()
}

func (g *GeneralMiddleware) GetNamespaceByName(name string) objects.NamespaceDto {
	return g.endpoint.GetNamespaceByName(name)
}

func (g *GeneralMiddleware) UpdateNamespaceByName(name string) bool {
	return g.endpoint.UpdateNamespaceByName(name)
}

func (g *GeneralMiddleware) DeleteNamespaceByName(name string) bool {
	return g.endpoint.DeleteNamespaceByName(name)
}
