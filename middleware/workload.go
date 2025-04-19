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

func (w *WorkloadMiddleware) GetPods() []objects.PodDto {
	return w.endpoint.GetPods()
}

func (w *WorkloadMiddleware) GetPodByName(name string) objects.PodDto {
	return w.endpoint.GetPodByName(name)
}

func (w *WorkloadMiddleware) UpdatePodByName(name string) bool {
	return w.endpoint.UpdatePodByName(name)
}

func (w *WorkloadMiddleware) RestartPodByName(name string) bool {
	return w.endpoint.RestartPodByName(name)
}

func (w *WorkloadMiddleware) GetDeployments() []objects.DeploymentDto {
	return w.endpoint.GetDeployments()
}

func (w *WorkloadMiddleware) GetDeploymentByName(name string) objects.DeploymentDto {
	return w.endpoint.GetDeploymentByName(name)
}

func (w *WorkloadMiddleware) UpdateDeploymentByName(name string) bool {
	return w.endpoint.UpdateDeploymentByName(name)
}

func (w *WorkloadMiddleware) DeleteDeploymentByName(name string) bool {
	return w.endpoint.DeleteDeploymentByName(name)
}
