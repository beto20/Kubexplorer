package usecase

import "Kubessistant/backend/objects"

type IServiceUseCase interface {
	GetServices() []objects.ServiceDto
	GetServiceByName(name string) objects.ServiceDto
	UpdateServiceByName(name string) bool
	DeleteServiceByName(name string) bool
}

type serviceImpl struct {
	object objects.IServiceObject
}

func NewServiceUseCase(object objects.IServiceObject) IServiceUseCase {
	return &serviceImpl{object: object}
}

func (s *serviceImpl) GetServices() []objects.ServiceDto {
	return s.object.GetServicesMock()
}

func (s *serviceImpl) GetServiceByName(name string) objects.ServiceDto {
	return s.object.GetServiceByName(name)
}

func (s *serviceImpl) UpdateServiceByName(name string) bool {
	return s.object.UpdateServiceByName(name)
}

func (s *serviceImpl) DeleteServiceByName(name string) bool {
	return s.object.DeleteServiceByName(name)
}
