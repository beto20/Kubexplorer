package endpoint

import (
	"Kubessistant/backend/model"
	"Kubessistant/backend/usecase"
)

type IWorkloadEndpoint interface {
	GetPods() ([]model.PodDto, error)
	//GetDeployments() []objects.DeploymentDto
	GetPod(name string) (model.PodDto, error)
	//GetDeploymentByName(name string) objects.DeploymentDto
	UpdatePod(name string) error
	UpdateDeployment(name string) bool
	RestartPod(name string) error
	DeleteDeployment(name string) bool
}

type workloadEndpoint struct {
	podUseCase        usecase.PodUseCase
	deploymentUseCase usecase.IDeploymentUseCase
}

func NewWorkloadEndpoint(podUseCase usecase.PodUseCase, deploymentUseCase usecase.IDeploymentUseCase) IWorkloadEndpoint {
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

//func (we *workloadEndpoint) GetDeployments() []objects.DeploymentDto {
//	return we.deploymentUseCase.GetAllDeployments()
//}
//
//func (we *workloadEndpoint) GetDeploymentByName(name string) objects.DeploymentDto {
//	return we.deploymentUseCase.GetDeploymentByName(name)
//}

func (we *workloadEndpoint) UpdateDeployment(name string) bool {
	return we.deploymentUseCase.UpdateDeploymentByName(name)
}

func (we *workloadEndpoint) DeleteDeployment(name string) bool {
	return we.deploymentUseCase.DeleteDeploymentByName(name)
}
