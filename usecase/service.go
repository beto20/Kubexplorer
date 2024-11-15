package usecase

import "github.com/beto20/kubessitant/objects"

type IServiceUseCase interface {
	GetAll(namespace string)
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

func (s *serviceImpl) GetAll(namespace string) {
	s.object.GetServicesMock(namespace)
}

func (s *serviceImpl) GetDetailsById() {
}

func (s *serviceImpl) DeleteOneById() {
}

func (s *serviceImpl) EditOneById() {
}
