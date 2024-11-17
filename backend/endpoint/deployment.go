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

func NewDeploymentEndpoint(useCase usecase.IDeploymentUseCase) *DeploymentEndpoint {
	return &DeploymentEndpoint{
		useCase: useCase,
	}
}

func (de *DeploymentEndpoint) GetDeployments(namespace string) []objects.DeploymentDto {
	return de.useCase.GetAllDeployments(namespace)

	//for _, d := range deployments {
	//	fmt.Printf("deployment name: %s namespace: %s status: %s age: %s\n",
	//		d.Name,
	//		d.Namespace,
	//		d.Status,
	//		d.Age,
	//	)
	//}
}
