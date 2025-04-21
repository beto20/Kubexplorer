package endpoint

import (
	"Kubessistant/backend/model"
	"Kubessistant/backend/usecase"
)

type IWorkloadEndpoint interface {
	GetPods() ([]model.PodDto, error)
	GetDeployments() ([]model.DeploymentDto, error)
	GetPod(name string) (model.PodDto, error)
	GetDeployment(name string) (model.DeploymentDto, error)
	UpdatePod(name string) error
	UpdateDeployment(name string) error
	RestartPod(name string) error
	DeleteDeployment(name string) error
}

type WorkloadEndpoint struct {
	podUseCase        usecase.PodUseCase
	deploymentUseCase usecase.DeploymentUseCase
}

func NewWorkloadEndpoint(podUseCase usecase.PodUseCase, deploymentUseCase usecase.DeploymentUseCase) IWorkloadEndpoint {
	return &WorkloadEndpoint{podUseCase: podUseCase, deploymentUseCase: deploymentUseCase}
}

func (we *WorkloadEndpoint) GetPods() ([]model.PodDto, error) {
	return we.podUseCase.GetAllPods()
}

func (we *WorkloadEndpoint) GetPod(name string) (model.PodDto, error) {
	return we.podUseCase.GetPod(name)
}

func (we *WorkloadEndpoint) UpdatePod(name string) error {
	return we.podUseCase.UpdatePod(name)
}

func (we *WorkloadEndpoint) RestartPod(name string) error {
	return we.podUseCase.DeletePod(name)
}

func (we *WorkloadEndpoint) GetDeployments() ([]model.DeploymentDto, error) {
	return we.deploymentUseCase.GetAllDeployments()
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
