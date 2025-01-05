package usecase

import "Kubessistant/backend/objects"

type IServiceUseCase interface {
	GetServices(namespace string) []objects.ServiceDto
	GetServiceByName(name string) objects.ServiceDto
}

type serviceImpl struct {
	object objects.IServiceObject
}

func NewServiceUseCase(object objects.IServiceObject) IServiceUseCase {
	return &serviceImpl{object: object}
}

func (s *serviceImpl) GetServices(namespace string) []objects.ServiceDto {
	return s.object.GetServicesMock(namespace)
}

func (s *serviceImpl) GetServiceByName(name string) objects.ServiceDto {
	return s.object.GetServiceByName(name)
}
