package usecase

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
)

type ServiceUseCase interface {
	GetServices(namespace string) ([]model.ServiceDto, error)
	GetService(name string, namespace string) (model.ServiceDto, error)
	UpdateService(name string, namespace string, dto model.ServiceDto) error
	DeleteService(name string, namespace string) error
}

type serviceUseCase struct {
	client kubeclient.ServiceClient
}

func NewServiceUseCase(client kubeclient.ServiceClient) ServiceUseCase {
	return &serviceUseCase{client: client}
}

func (s *serviceUseCase) GetServices(namespace string) ([]model.ServiceDto, error) {
	return s.client.GetServices(namespace)
}

func (s *serviceUseCase) GetService(name string, namespace string) (model.ServiceDto, error) {
	return s.client.GetService(name, namespace)
}

func (s *serviceUseCase) UpdateService(name string, namespace string, dto model.ServiceDto) error {
	return s.client.UpdateService(name, namespace, dto)
}

func (s *serviceUseCase) DeleteService(name string, namespace string) error {
	return s.client.DeleteService(name, namespace)
}
