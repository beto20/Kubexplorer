package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type IServiceEndpoint interface {
	GetServices(namespace string) []objects.ServiceDto
	GetIngresses(namespace string) []objects.IngressDto
}

type ServiceEndpoint struct {
	useCase usecase.IServiceUseCase
}

func NewServiceEndpoint(useCase usecase.IServiceUseCase) IServiceEndpoint {
	return &ServiceEndpoint{useCase: useCase}
}

func (se *ServiceEndpoint) GetServices(namespace string) []objects.ServiceDto {
	return se.useCase.GetServices(namespace)
}

func (se *ServiceEndpoint) GetIngresses(namespace string) []objects.IngressDto {
	return se.useCase.GetIngresses(namespace)
}
