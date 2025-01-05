package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type IPodEndpoint interface {
	GetPods(namespace string) []objects.PodDto
}

type podEndpoint struct {
	useCase usecase.IPodUseCase
}

func NewPodEndpoint(useCase usecase.IPodUseCase) IPodEndpoint {
	return &podEndpoint{useCase: useCase}
}

func (pe *podEndpoint) GetPods(namespace string) []objects.PodDto {
	return pe.useCase.GetAllPods(namespace)
}

func (pe *podEndpoint) GetPodById(namespace string, id string) {

}

func (pe *podEndpoint) DeletePodById(namespace string, id string) {

}

func (pe *podEndpoint) UpdatePodById(namespace string, id string) {

}
