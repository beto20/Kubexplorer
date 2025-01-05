package middleware

import (
	"Kubessistant/backend/endpoint"
	"Kubessistant/backend/objects"
)

type WorkloadMiddleware struct {
	endpoint endpoint.IWorkloadEndpoint
}

func NewWorkloadMiddleware(endpoint endpoint.IWorkloadEndpoint) *WorkloadMiddleware {
	return &WorkloadMiddleware{endpoint: endpoint}
}

func (d *WorkloadMiddleware) GetPods(namespace string) []objects.PodDto {
	return d.endpoint.GetPods(namespace)
}

func (d *WorkloadMiddleware) GetDeployments(namespace string) []objects.DeploymentDto {
	return d.endpoint.GetDeployments(namespace)
}
