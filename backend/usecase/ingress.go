package usecase

import "Kubessistant/backend/objects"

type IIngressUseCase interface {
	GetIngresses(namespace string) []objects.IngressDto
	GetIngressByName(name string) objects.IngressDto
	UpdateIngressByName(name string) bool
	DeleteIngressByName(name string) bool
}

type ingressImpl struct {
	object objects.IIngressObject
}

func NewIngressUseCase(object objects.IIngressObject) IIngressUseCase {
	return &ingressImpl{object: object}
}

func (i *ingressImpl) GetIngresses(namespace string) []objects.IngressDto {
	return i.object.GetIngresses(namespace)
}

func (i *ingressImpl) GetIngressByName(name string) objects.IngressDto {
	return i.object.GetIngressByName(name)
}

func (i *ingressImpl) UpdateIngressByName(name string) bool {
	return i.object.UpdateIngressByName(name)
}

func (i *ingressImpl) DeleteIngressByName(name string) bool {
	return i.object.DeleteIngressByName(name)
}
