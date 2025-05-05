package usecase

import (
	"Kubessistant/backend/kubeclient"
	"Kubessistant/backend/model"
)

type IngressUseCase interface {
	GetIngresses() ([]model.IngressDto, error)
	GetIngress(name string) (model.IngressDto, error)
	UpdateIngress(name string) error
	DeleteIngress(name string) error
}

type ingressUseCase struct {
	client kubeclient.IngressClient
}

func NewIngressUseCase(client kubeclient.IngressClient) IngressUseCase {
	return &ingressUseCase{client: client}
}

func (i *ingressUseCase) GetIngresses() ([]model.IngressDto, error) {
	return i.client.GetIngresses()
}

func (i *ingressUseCase) GetIngress(name string) (model.IngressDto, error) {
	return i.client.GetIngress(name)
}

func (i *ingressUseCase) UpdateIngress(name string) error {
	return i.client.UpdateIngress(name)
}

func (i *ingressUseCase) DeleteIngress(name string) error {
	return i.client.DeleteIngress(name)
}
