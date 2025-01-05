package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type INetworkEndpoint interface {
	GetServices(namespace string) []objects.ServiceDto
	GetIngresses(namespace string) []objects.IngressDto
}

type networkEndpoint struct {
	useCase usecase.IServiceUseCase
}

func NewNetworkEndpoint(useCase usecase.IServiceUseCase) INetworkEndpoint {
	return &networkEndpoint{useCase: useCase}
}

func (se *networkEndpoint) GetServices(namespace string) []objects.ServiceDto {
	return se.useCase.GetServices(namespace)
}

func (se *networkEndpoint) GetIngresses(namespace string) []objects.IngressDto {
	return se.useCase.GetIngresses(namespace)
}
