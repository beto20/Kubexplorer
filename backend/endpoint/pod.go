package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type IPodEndpoint interface {
	GetPods(namespace string) []objects.PodDto
}

type PodEndpoint struct {
	useCase usecase.IPodUseCase
}

func NewPodEndpoint(useCase usecase.IPodUseCase) IPodEndpoint {
	return &PodEndpoint{useCase: useCase}
}

func (pe *PodEndpoint) GetPods(namespace string) []objects.PodDto {
	return pe.useCase.GetAllPods(namespace)
}

func (pe *PodEndpoint) GetPodById(namespace string, id string) {

}

func (pe *PodEndpoint) DeletePodById(namespace string, id string) {

}

func (pe *PodEndpoint) UpdatePodById(namespace string, id string) {

}
