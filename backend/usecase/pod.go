package usecase

import (
	"Kubessistant/backend/kubeclient"
	"Kubessistant/backend/model"
)

type PodUseCase interface {
	GetAllPods() ([]model.PodDto, error)
	GetPod(name string) (model.PodDto, error)
	UpdatePod(name string) error
	DeletePod(name string) error
}

type podUseCase struct {
	client kubeclient.PodClient
}

func NewPodUseCase(client kubeclient.PodClient) PodUseCase {
	return &podUseCase{client: client}
}

func (p *podUseCase) GetAllPods() ([]model.PodDto, error) {
	return p.client.GetPodsMock()
}

func (p *podUseCase) GetPod(name string) (model.PodDto, error) {
	return p.client.GetPod(name)
}

func (p *podUseCase) UpdatePod(name string) error {
	return p.client.UpdatePod(name)
}

func (p *podUseCase) DeletePod(name string) error {
	return p.client.DeletePod(name)
}
