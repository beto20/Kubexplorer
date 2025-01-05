package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type IDeploymentEndpoint interface {
	GetDeployments(namespace string) []objects.DeploymentDto
}

type deploymentEndpoint struct {
	useCase usecase.IDeploymentUseCase
}

func NewDeploymentEndpoint(useCase usecase.IDeploymentUseCase) IDeploymentEndpoint {
	return &deploymentEndpoint{
		useCase: useCase,
	}
}

func (de *deploymentEndpoint) GetDeployments(namespace string) []objects.DeploymentDto {
	return de.useCase.GetAllDeployments(namespace)
}
