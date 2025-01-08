package middleware

import (
	"Kubessistant/backend/endpoint"
	"Kubessistant/backend/objects"
)

type NetworkMiddleware struct {
	endpoint endpoint.INetworkEndpoint
}

func NewNetworkMiddleware(endpoint endpoint.INetworkEndpoint) *NetworkMiddleware {
	return &NetworkMiddleware{endpoint: endpoint}
}

func (n *NetworkMiddleware) GetServices() []objects.ServiceDto {
	return n.endpoint.GetServices("")
}

func (n *NetworkMiddleware) GetServiceByName(name string) objects.ServiceDto {
	return n.endpoint.GetServiceByName(name)
}

func (n *NetworkMiddleware) UpdateServiceByName(name string) bool {
	return n.endpoint.UpdateServiceByName(name)
}

func (n *NetworkMiddleware) DeleteServiceByName(name string) bool {
	return n.endpoint.DeleteServiceByName(name)
}

func (n *NetworkMiddleware) GetIngresses() []objects.IngressDto {
	return n.endpoint.GetIngresses("")
}

func (n *NetworkMiddleware) GetIngressByName(name string) objects.IngressDto {
	return n.endpoint.GetIngressByName(name)
}

func (n *NetworkMiddleware) UpdateIngressByName(name string) bool {
	return n.endpoint.UpdateIngressByName(name)
}

func (n *NetworkMiddleware) DeleteIngressByName(name string) bool {
	return n.endpoint.DeleteIngressByName(name)
}
