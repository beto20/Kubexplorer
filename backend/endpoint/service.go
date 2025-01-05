package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type IServiceEndpoint interface {
	GetServices(namespace string) []objects.ServiceDto
	GetIngresses(namespace string) []objects.IngressDto
}

type serviceEndpoint struct {
	useCase usecase.IServiceUseCase
}

func NewServiceEndpoint(useCase usecase.IServiceUseCase) IServiceEndpoint {
	return &serviceEndpoint{useCase: useCase}
}

func (se *serviceEndpoint) GetServices(namespace string) []objects.ServiceDto {
	return se.useCase.GetServices(namespace)
}

func (se *serviceEndpoint) GetIngresses(namespace string) []objects.IngressDto {
	return se.useCase.GetIngresses(namespace)
}
