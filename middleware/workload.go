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

func (d *WorkloadMiddleware) GetPodByName(name string) objects.PodDto {
	return d.endpoint.GetPodByName(name)
}

func (d *WorkloadMiddleware) UpdatePodByName(name string) bool {
	return d.endpoint.UpdatePodByName(name)
}

func (d *WorkloadMiddleware) RestartPodByName(name string) bool {
	return d.endpoint.RestartPodByName(name)
}

func (d *WorkloadMiddleware) GetDeployments(namespace string) []objects.DeploymentDto {
	return d.endpoint.GetDeployments(namespace)
}

func (d *WorkloadMiddleware) GetDeploymentByName(name string) objects.DeploymentDto {
	return d.endpoint.GetDeploymentByName(name)
}

func (d *WorkloadMiddleware) UpdateDeploymentByName(name string) bool {
	return d.endpoint.UpdateDeploymentByName(name)
}

func (d *WorkloadMiddleware) DeleteDeploymentByName(name string) bool {
	return d.endpoint.DeleteDeploymentByName(name)
}
