package endpoint

import "github.com/beto20/kubessitant/usecase"

type IServiceEndpoint interface {
	getServices()
}

type ServiceEndpoint struct {
	useCase usecase.Service
}

func NewServiceEndpoint(useCase usecase.Service) *ServiceEndpoint {
	return &ServiceEndpoint{useCase: useCase}
}

func (se ServiceEndpoint) getServices() {
	se.useCase.GetAll()
}
