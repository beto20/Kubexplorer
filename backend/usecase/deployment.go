package usecase

import (
	"Kubessistant/backend/kubeclient2"
	"Kubessistant/backend/model"
)

type DeploymentUseCase interface {
	GetAllDeployments() ([]model.DeploymentDto, error)
	GetDeployment(name string) (model.DeploymentDto, error)
	UpdateDeployment(name string) error
	DeleteDeployment(name string) error
}

type deploymentUseCase struct {
	client kubeclient2.DeploymentClient
}

func NewDeploymentUseCase(client kubeclient2.DeploymentClient) DeploymentUseCase {
	return &deploymentUseCase{client: client}
}

func (d *deploymentUseCase) GetAllDeployments() ([]model.DeploymentDto, error) {
	return d.client.GetDeploymentsMock()
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
