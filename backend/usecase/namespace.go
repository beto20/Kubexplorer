package usecase

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/service"
)

type NamespaceUseCase interface {
	GetNamespaces() ([]model.NamespaceDto, error)
	GetNamespace(name string) (model.NamespaceDto, error)
	UpdateNamespace(name string, dto model.NamespaceDto) error
	DeleteNamespace(name string) error
	ExportNamespaceObjects(namespace string, directory string) error
}

type namespaceUseCase struct {
	client  kubeclient.NamespaceClient
	service service.NamespaceService
}

func NewNamespaceUseCase(client kubeclient.NamespaceClient) NamespaceUseCase {
	return &namespaceUseCase{client: client}
}

func (n *namespaceUseCase) GetNamespaces() ([]model.NamespaceDto, error) {
	return n.client.GetNamespaces()
}
func (n *namespaceUseCase) GetNamespace(name string) (model.NamespaceDto, error) {
	return n.client.GetNamespace(name)
}

func (n *namespaceUseCase) UpdateNamespace(name string, dto model.NamespaceDto) error {
	return n.client.UpdateNamespace(name, dto)
}

func (n *namespaceUseCase) DeleteNamespace(name string) error {
	return n.client.DeleteNamespace(name)
}

func (n *namespaceUseCase) ExportNamespaceObjects(namespace string, directory string) error {
	return n.service.ExportObjects(namespace, directory)
}
