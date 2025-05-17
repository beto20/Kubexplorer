package usecase

import (
	"Kubessistant/backend/kubeclient"
	"Kubessistant/backend/model"
)

type DeploymentUseCase interface {
	GetDeployments() ([]model.DeploymentDto, error)
	GetDeployment(name string) (model.DeploymentDto, error)
	UpdateDeployment(name string) error
	DeleteDeployment(name string) error
}

type deploymentUseCase struct {
	client kubeclient.DeploymentClient
}

func NewDeploymentUseCase(client kubeclient.DeploymentClient) DeploymentUseCase {
	return &deploymentUseCase{client: client}
}

func (d *deploymentUseCase) GetDeployments() ([]model.DeploymentDto, error) {
	return d.client.GetDeployments()
}

func (d *deploymentUseCase) GetDeployment(name string) (model.DeploymentDto, error) {
	return d.client.GetDeployment(name)
}

func (d *deploymentUseCase) UpdateDeployment(name string) error {
	return d.client.UpdateDeployment(name)
}

func (d *deploymentUseCase) DeleteDeployment(name string) error {
	return d.client.DeleteDeployment(name)
}
