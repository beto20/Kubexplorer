package endpoint

import (
	"Kubessistant/backend/usecase"
	"fmt"
)

type IServiceEndpoint interface {
	GetServices(namespace string)
}

type ServiceEndpoint struct {
	useCase usecase.IServiceUseCase
}

func NewServiceEndpoint(useCase usecase.IServiceUseCase) *ServiceEndpoint {
	return &ServiceEndpoint{useCase: useCase}
}

func (se ServiceEndpoint) GetServices(namespace string) {
	se.useCase.GetAll(namespace)
	fmt.Printf("NS: %s", namespace)
}
