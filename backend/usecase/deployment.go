package usecase

import (
	"Kubessistant/backend/kubeclient"
	"Kubessistant/backend/model"
)

type DeploymentUseCase interface {
	GetDeployments() ([]model.DeploymentDto, error)
	GetDeployment(name string, namespace string) (model.DeploymentDto, error)
	UpdateDeployment(name string, namespace string, dto model.DeploymentDto) error
	DeleteDeployment(name string, namespace string) error
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

func (d *deploymentUseCase) GetDeployment(name string, namespace string) (model.DeploymentDto, error) {
	return d.client.GetDeployment(name, namespace)
}

func (d *deploymentUseCase) UpdateDeployment(name string, namespace string, dto model.DeploymentDto) error {
	return d.client.UpdateDeployment(name, namespace, dto)
}

func (d *deploymentUseCase) DeleteDeployment(name string, namespace string) error {
	return d.client.DeleteDeployment(name, namespace)
}
