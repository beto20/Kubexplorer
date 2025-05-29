package endpoint

import (
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/usecase"
)

type WorkloadEndpoint struct {
	podUseCase        usecase.PodUseCase
	deploymentUseCase usecase.DeploymentUseCase
	resourceUseCase   usecase.ResourceUseCase
}

func NewWorkloadEndpoint(podUseCase usecase.PodUseCase, deploymentUseCase usecase.DeploymentUseCase) *WorkloadEndpoint {
	return &WorkloadEndpoint{podUseCase: podUseCase, deploymentUseCase: deploymentUseCase}
}

func (we *WorkloadEndpoint) GetPods() ([]model.PodDto, error) {
	return we.podUseCase.GetPods()
}

func (we *WorkloadEndpoint) GetPod(name string, namespace string) (model.PodDto, error) {
	return we.podUseCase.GetPod(name, namespace)
}

func (we *WorkloadEndpoint) UpdatePod(name string, namespace string, dto model.PodDto) error {
	return we.podUseCase.UpdatePod(name, namespace, dto)
}

func (we *WorkloadEndpoint) RestartPod(name string, namespace string) error {
	return we.podUseCase.RestartPod(name, namespace)
}

func (we *WorkloadEndpoint) GetDeployments() ([]model.DeploymentDto, error) {
	return we.deploymentUseCase.GetDeployments()
}

func (we *WorkloadEndpoint) GetDeployment(name string, namespace string) (model.DeploymentDto, error) {
	return we.deploymentUseCase.GetDeployment(name, namespace)
}

func (we *WorkloadEndpoint) UpdateDeployment(name string, namespace string, dto model.DeploymentDto) error {
	return we.deploymentUseCase.UpdateDeployment(name, namespace, dto)
}

func (we *WorkloadEndpoint) DeleteDeployment(name string, namespace string) error {
	return we.deploymentUseCase.DeleteDeployment(name, namespace)
}

func (we *WorkloadEndpoint) ResourceTuning(namespace string) {
	we.resourceUseCase.Invoke(namespace)
}
