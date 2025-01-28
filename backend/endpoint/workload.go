package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type IWorkloadEndpoint interface {
	GetPods() []objects.PodDto
	GetDeployments() []objects.DeploymentDto
	GetPodByName(name string) objects.PodDto
	GetDeploymentByName(name string) objects.DeploymentDto
	UpdatePodByName(name string) bool
	UpdateDeploymentByName(name string) bool
	RestartPodByName(name string) bool
	DeleteDeploymentByName(name string) bool
}

type workloadEndpoint struct {
	podUseCase        usecase.IPodUseCase
	deploymentUseCase usecase.IDeploymentUseCase
}

func NewWorkloadEndpoint(podUseCase usecase.IPodUseCase, deploymentUseCase usecase.IDeploymentUseCase) IWorkloadEndpoint {
	return &workloadEndpoint{podUseCase: podUseCase, deploymentUseCase: deploymentUseCase}
}

func (we *workloadEndpoint) GetPods() []objects.PodDto {
	return we.podUseCase.GetAllPods()
}

func (we *workloadEndpoint) GetPodByName(name string) objects.PodDto {
	return we.podUseCase.GetPodByName(name)
}

func (we *workloadEndpoint) UpdatePodByName(name string) bool {
	return we.podUseCase.UpdatePodByName(name)
}

func (we *workloadEndpoint) RestartPodByName(name string) bool {
	return we.podUseCase.DeletePodByName(name)
}

func (we *workloadEndpoint) GetDeployments() []objects.DeploymentDto {
	return we.deploymentUseCase.GetAllDeployments()
}

func (we *workloadEndpoint) GetDeploymentByName(name string) objects.DeploymentDto {
	return we.deploymentUseCase.GetDeploymentByName(name)
}

func (we *workloadEndpoint) UpdateDeploymentByName(name string) bool {
	return we.deploymentUseCase.UpdateDeploymentByName(name)
}

func (we *workloadEndpoint) DeleteDeploymentByName(name string) bool {
	return we.deploymentUseCase.DeleteDeploymentByName(name)
}
