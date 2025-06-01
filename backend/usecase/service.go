package usecase

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/service"
)

type ServiceUseCase interface {
	GetServices(namespace string) ([]model.ServiceDto, error)
	GetService(name string, namespace string) (model.ServiceDto, error)
	UpdateService(name string, namespace string, dto model.ServiceDto) error
	DeleteService(name string, namespace string) error
	TroubleshootService(name string, namespace string)
}

type serviceUseCase struct {
	client  kubeclient.ServiceClient
	service service.Troubleshooting
}

func NewServiceUseCase(client kubeclient.ServiceClient, service *service.Troubleshooting) ServiceUseCase {
	return &serviceUseCase{client: client, service: *service}
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

func (s *serviceUseCase) TroubleshootService(name string, namespace string) {
	s.service.Analyse(name, namespace, service.Service)
}
