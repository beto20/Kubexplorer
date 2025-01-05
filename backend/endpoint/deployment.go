package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type IDeploymentEndpoint interface {
	GetDeployments(namespace string) []objects.DeploymentDto
}

type DeploymentEndpoint struct {
	useCase usecase.IDeploymentUseCase
}

func NewDeploymentEndpoint(useCase usecase.IDeploymentUseCase) IDeploymentEndpoint {
	return &DeploymentEndpoint{
		useCase: useCase,
	}
}

func (de *DeploymentEndpoint) GetDeployments(namespace string) []objects.DeploymentDto {
	return de.useCase.GetAllDeployments(namespace)
}
