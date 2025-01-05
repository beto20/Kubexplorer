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

func (s *NetworkMiddleware) GetServices() []objects.ServiceDto {
	return s.endpoint.GetServices("")
}

func (s *NetworkMiddleware) GetIngresses() []objects.IngressDto {
	return s.endpoint.GetIngresses("")
}

func (s *NetworkMiddleware) GetServiceByName(name string) objects.ServiceDto {
	return s.endpoint.GetServiceByName(name)
}

func (s *NetworkMiddleware) GetIngressByName(name string) objects.IngressDto {
	return s.endpoint.GetIngressByName(name)
}
