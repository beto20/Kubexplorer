package usecase

import (
	"Kubessistant/backend/objects"
)

type IPodUseCase interface {
	GetAllPods(namespace string) []objects.PodDto
}

type podImpl struct {
	object objects.IPodObject
}

func NewPodUseCase(object objects.IPodObject) IPodUseCase {
	return &podImpl{object: object}
}

func (p *podImpl) GetAllPods(namespace string) []objects.PodDto {
	return p.object.GetPodsMock(namespace)
}

func (p *podImpl) GetDetailsById() {
}

func (p *podImpl) DeleteOneById() {
}

func (p *podImpl) EditOneById() {
}
