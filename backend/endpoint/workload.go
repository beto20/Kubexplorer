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

type workloadEndpoint struct {
	podUseCase        usecase.PodUseCase
	deploymentUseCase usecase.DeploymentUseCase
}

func NewWorkloadEndpoint(podUseCase usecase.PodUseCase, deploymentUseCase usecase.DeploymentUseCase) IWorkloadEndpoint {
	return &workloadEndpoint{podUseCase: podUseCase, deploymentUseCase: deploymentUseCase}
}

func (we *workloadEndpoint) GetPods() ([]model.PodDto, error) {
	return we.podUseCase.GetAllPods()
}

func (we *workloadEndpoint) GetPod(name string) (model.PodDto, error) {
	return we.podUseCase.GetPod(name)
}

func (we *workloadEndpoint) UpdatePod(name string) error {
	return we.podUseCase.UpdatePod(name)
}

func (we *workloadEndpoint) RestartPod(name string) error {
	return we.podUseCase.DeletePod(name)
}

func (we *workloadEndpoint) GetDeployments() ([]model.DeploymentDto, error) {
	return we.deploymentUseCase.GetAllDeployments()
}

func (we *workloadEndpoint) GetDeployment(name string) (model.DeploymentDto, error) {
	return we.deploymentUseCase.GetDeployment(name)
}

func (we *workloadEndpoint) UpdateDeployment(name string) error {
	return we.deploymentUseCase.UpdateDeployment(name)
}

func (we *workloadEndpoint) DeleteDeployment(name string) error {
	return we.deploymentUseCase.DeleteDeployment(name)
}
