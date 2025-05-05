package endpoint

import (
	"Kubessistant/backend/model"
	"Kubessistant/backend/usecase"
)

type WorkloadEndpoint struct {
	podUseCase        usecase.PodUseCase
	deploymentUseCase usecase.DeploymentUseCase
}

func NewWorkloadEndpoint(podUseCase usecase.PodUseCase, deploymentUseCase usecase.DeploymentUseCase) *WorkloadEndpoint {
	return &WorkloadEndpoint{podUseCase: podUseCase, deploymentUseCase: deploymentUseCase}
}

func (we *WorkloadEndpoint) GetPods() ([]model.PodDto, error) {
	return we.podUseCase.GetPods()
}

func (we *WorkloadEndpoint) GetPod(name string) (model.PodDto, error) {
	return we.podUseCase.GetPod(name)
}

func (we *WorkloadEndpoint) UpdatePod(name string) error {
	return we.podUseCase.UpdatePod(name)
}

func (we *WorkloadEndpoint) RestartPod(name string) error {
	return we.podUseCase.RestartPod(name)
}

func (we *WorkloadEndpoint) GetDeployments() ([]model.DeploymentDto, error) {
	return we.deploymentUseCase.GetDeployments()
}

func (we *WorkloadEndpoint) GetDeployment(name string) (model.DeploymentDto, error) {
	return we.deploymentUseCase.GetDeployment(name)
}

func (we *WorkloadEndpoint) UpdateDeployment(name string) error {
	return we.deploymentUseCase.UpdateDeployment(name)
}

func (we *WorkloadEndpoint) DeleteDeployment(name string) error {
	return we.deploymentUseCase.DeleteDeployment(name)
}
