package usecase

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
)

type IngressUseCase interface {
	GetIngresses(namespace string) ([]model.IngressDto, error)
	GetIngress(name string, namespace string) (model.IngressDto, error)
	UpdateIngress(name string, namespace string, dto model.IngressDto) error
	DeleteIngress(name string, namespace string) error
}

type ingressUseCase struct {
	client kubeclient.IngressClient
}

func NewIngressUseCase(client kubeclient.IngressClient) IngressUseCase {
	return &ingressUseCase{client: client}
}

func (i *ingressUseCase) GetIngresses(namespace string) ([]model.IngressDto, error) {
	return i.client.GetIngresses()
}

func (i *ingressUseCase) GetIngress(name string, namespace string) (model.IngressDto, error) {
	return i.client.GetIngress(name, namespace)
}

func (i *ingressUseCase) UpdateIngress(name string, namespace string, dto model.IngressDto) error {
	return i.client.UpdateIngress(name, namespace, dto)
}

func (i *ingressUseCase) DeleteIngress(name string, namespace string) error {
	return i.client.DeleteIngress(name, namespace)
}
