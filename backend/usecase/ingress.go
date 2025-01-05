package usecase

import "Kubessistant/backend/objects"

type IIngressUseCase interface {
	GetIngresses(namespace string) []objects.IngressDto
	GetIngressByName(name string) objects.IngressDto
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
