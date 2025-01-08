package usecase

import (
	"Kubessistant/backend/objects"
)

type IPodUseCase interface {
	GetAllPods(namespace string) []objects.PodDto
	GetPodByName(name string) objects.PodDto
	UpdatePodByName(name string) bool
	DeletePodByName(name string) bool
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

func (p *podImpl) GetPodByName(name string) objects.PodDto {
	return p.object.GetPodByName(name)
}

func (p *podImpl) UpdatePodByName(name string) bool {
	return p.object.UpdatePodByName(name)
}

func (p *podImpl) DeletePodByName(name string) bool {
	return p.object.DeletePodByName(name)
}
