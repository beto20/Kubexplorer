package middleware

import (
	"Kubessistant/backend/endpoint"
	"Kubessistant/backend/model"
)

type WorkloadMiddleware struct {
	endpoint endpoint.IWorkloadEndpoint
}

func NewWorkloadMiddleware(endpoint endpoint.IWorkloadEndpoint) *WorkloadMiddleware {
	return &WorkloadMiddleware{endpoint: endpoint}
}

func (w *WorkloadMiddleware) GetPods() ([]model.PodDto, error) {
	return w.endpoint.GetPods()
}

func (w *WorkloadMiddleware) GetPod(name string) (model.PodDto, error) {
	return w.endpoint.GetPod(name)
}

func (w *WorkloadMiddleware) UpdatePod(name string) error {
	return w.endpoint.UpdatePod(name)
}

func (w *WorkloadMiddleware) RestartPod(name string) error {
	return w.endpoint.RestartPod(name)
}

func (w *WorkloadMiddleware) GetDeployments() ([]model.DeploymentDto, error) {
	return w.endpoint.GetDeployments()
}

func (w *WorkloadMiddleware) GetDeployment(name string) (model.DeploymentDto, error) {
	return w.endpoint.GetDeployment(name)
}

func (w *WorkloadMiddleware) UpdateDeployment(name string) error {
	return w.endpoint.UpdateDeployment(name)
}

func (w *WorkloadMiddleware) DeleteDeployment(name string) error {
	return w.endpoint.DeleteDeployment(name)
}
