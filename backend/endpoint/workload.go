package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type IWorkloadEndpoint interface {
	GetPods(namespace string) []objects.PodDto
	GetDeployments(namespace string) []objects.DeploymentDto
	GetPodByName(name string) objects.PodDto
	GetDeploymentByName(name string) objects.DeploymentDto
}

type workloadEndpoint struct {
	podUseCase        usecase.IPodUseCase
	deploymentUseCase usecase.IDeploymentUseCase
}

func NewWorkloadEndpoint(podUseCase usecase.IPodUseCase, deploymentUseCase usecase.IDeploymentUseCase) IWorkloadEndpoint {
	return &workloadEndpoint{podUseCase: podUseCase, deploymentUseCase: deploymentUseCase}
}

func (we *workloadEndpoint) GetPods(namespace string) []objects.PodDto {
	return we.podUseCase.GetAllPods(namespace)
}

func (we *workloadEndpoint) GetPodByName(name string) objects.PodDto {
	return we.podUseCase.GetPodByName(name)
}

func (we *workloadEndpoint) GetDeployments(namespace string) []objects.DeploymentDto {
	return we.deploymentUseCase.GetAllDeployments(namespace)
}

func (we *workloadEndpoint) GetDeploymentByName(name string) objects.DeploymentDto {
	return we.deploymentUseCase.GetDeploymentByName(name)
}
