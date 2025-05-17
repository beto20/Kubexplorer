package middleware

import (
	"Kubessistant/backend/endpoint"
	"Kubessistant/backend/kubeclient"
	"Kubessistant/backend/model"
	"Kubessistant/backend/usecase"
	"k8s.io/client-go/kubernetes"
)

type WorkloadMiddleware struct {
	endpoint endpoint.WorkloadEndpoint
}

func NewWorkloadMiddleware(endpoint *endpoint.WorkloadEndpoint) *WorkloadMiddleware {
	return &WorkloadMiddleware{endpoint: *endpoint}
}

func (w *WorkloadMiddleware) GetPods() []model.PodDto {
	x, _ := w.endpoint.GetPods()
	return x
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

func (w *WorkloadMiddleware) GetDeployments() []model.DeploymentDto {
	x, _ := w.endpoint.GetDeployments()
	return x
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

func BuildWorkload(client kubernetes.Interface) *WorkloadMiddleware {
	deploymentClient := kubeclient.NewDeployment(client)
	podClient := kubeclient.NewPod(client)

	deploymentUseCase := usecase.NewDeploymentUseCase(deploymentClient)
	podUseCase := usecase.NewPodUseCase(podClient)

	workloadEndpoint := endpoint.NewWorkloadEndpoint(podUseCase, deploymentUseCase)

	return NewWorkloadMiddleware(workloadEndpoint)
}
