package endpoint

import "github.com/beto20/kubessitant/usecase"

type ServiceEndpoint struct {
	usecase usecase.K8sObject
}

func NewServiceEndpoint(usecase usecase.K8sObject) *ServiceEndpoint {
	return &ServiceEndpoint{}
}

func (se ServiceEndpoint) getServices() {
	se.usecase.GetAll()
}
