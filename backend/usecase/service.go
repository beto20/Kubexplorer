package usecase

import (
	"Kubessistant/backend/kubeclient"
	"Kubessistant/backend/model"
)

type ServiceUseCase interface {
	GetServices() ([]model.ServiceDto, error)
	GetService(name string) (model.ServiceDto, error)
	UpdateService(name string) error
	DeleteService(name string) error
}

type serviceUseCase struct {
	client kubeclient.ServiceClient
}

func NewServiceUseCase(client kubeclient.ServiceClient) ServiceUseCase {
	return &serviceUseCase{client: client}
}

func (s *serviceUseCase) GetServices() ([]model.ServiceDto, error) {
	return s.client.GetServicesMock()
}

func (s *serviceUseCase) GetService(name string) (model.ServiceDto, error) {
	return s.client.GetService(name)
}

func (s *serviceUseCase) UpdateService(name string) error {
	return s.client.UpdateService(name)
}

func (s *serviceUseCase) DeleteService(name string) error {
	return s.client.DeleteService(name)
}
