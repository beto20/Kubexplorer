package usecase

import "Kubessistant/backend/objects"

type IServiceUseCase interface {
	GetServices(namespace string) []objects.ServiceDto
	GetIngresses(namespace string) []objects.IngressDto
	GetDetailsById()
	DeleteOneById()
	EditOneById()
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

func (s *serviceImpl) GetIngresses(namespace string) []objects.IngressDto {
	return s.object.GetIngresses(namespace)
}

func (s *serviceImpl) GetDetailsById() {
}

func (s *serviceImpl) DeleteOneById() {
}

func (s *serviceImpl) EditOneById() {
}
