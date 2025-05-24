package usecase

import (
	"Kubessistant/backend/kubeclient"
	"Kubessistant/backend/model"
)

type PodUseCase interface {
	GetPods() ([]model.PodDto, error)
	GetPod(name string, namespace string) (model.PodDto, error)
	UpdatePod(name string, namespace string, dto model.PodDto) error
	RestartPod(name string, namespace string) error
}

type podUseCase struct {
	client kubeclient.PodClient
}

func NewPodUseCase(client kubeclient.PodClient) PodUseCase {
	return &podUseCase{client: client}
}

func (p *podUseCase) GetPods() ([]model.PodDto, error) {
	return p.client.GetPods()
}

func (p *podUseCase) GetPod(name string, namespace string) (model.PodDto, error) {
	return p.client.GetPod(name, namespace)
}

func (p *podUseCase) UpdatePod(name string, namespace string, dto model.PodDto) error {
	return p.client.UpdatePod(name, namespace, dto)
}

func (p *podUseCase) RestartPod(name string, namespace string) error {
	return p.client.DeletePod(name, namespace)
}
