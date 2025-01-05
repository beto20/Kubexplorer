package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type IWorkloadEndpoint interface {
	GetPods(namespace string) []objects.PodDto
	GetDeployments(namespace string) []objects.DeploymentDto
}

type workloadEndpoint struct {
	podUseCase        usecase.IPodUseCase
	deploymentUseCase usecase.IDeploymentUseCase
}

func NewWorkloadEndpoint(podUseCase usecase.IPodUseCase, deploymentUseCase usecase.IDeploymentUseCase) IWorkloadEndpoint {
	return &workloadEndpoint{podUseCase: podUseCase, deploymentUseCase: deploymentUseCase}
}

func (pe *workloadEndpoint) GetPods(namespace string) []objects.PodDto {
	return pe.podUseCase.GetAllPods(namespace)
}

func (pe *workloadEndpoint) GetPodById(namespace string, id string) {

}

func (pe *workloadEndpoint) DeletePodById(namespace string, id string) {

}

func (pe *workloadEndpoint) UpdatePodById(namespace string, id string) {

}

func (pe *workloadEndpoint) GetDeployments(namespace string) []objects.DeploymentDto {
	return pe.deploymentUseCase.GetAllDeployments(namespace)
}
