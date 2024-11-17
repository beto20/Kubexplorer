package middleware

import (
	"Kubessistant/backend/endpoint"
	"Kubessistant/backend/objects"
)

type PodMiddleware struct {
	endpoint endpoint.IPodEndpoint
}

func NewPodMiddleware(endpoint endpoint.IPodEndpoint) *PodMiddleware {
	return &PodMiddleware{endpoint: endpoint}
}

func (d *PodMiddleware) GetPods(namespace string) []objects.PodDto {
	return d.endpoint.GetPods(namespace)
}
