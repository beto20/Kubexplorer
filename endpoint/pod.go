package endpoint

import "github.com/beto20/kubessitant/usecase"

type PodEndpoint struct {
	usecase usecase.K8sObject
}

func NewPodEndpoint(usecase usecase.K8sObject) *PodEndpoint {
	return &PodEndpoint{}
}

func (pe *PodEndpoint) getPods() {
	pe.usecase.GetAll()
}
