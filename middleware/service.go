package middleware

import (
	"Kubessistant/backend/endpoint"
	"Kubessistant/backend/objects"
)

type ServiceMiddleware struct {
	endpoint endpoint.IServiceEndpoint
}

func NewServiceMiddleware(endpoint endpoint.IServiceEndpoint) *ServiceMiddleware {
	return &ServiceMiddleware{endpoint: endpoint}
}

func (s *ServiceMiddleware) GetServices() []objects.ServiceDto {
	return s.endpoint.GetServices("")
}

func (s *ServiceMiddleware) GetIngresses() []objects.IngressDto {
	return s.endpoint.GetIngresses("")
}
