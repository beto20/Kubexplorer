package usecase

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/service"
)

type DeploymentUseCase interface {
	GetDeployments() ([]model.DeploymentDto, error)
	GetDeployment(name string, namespace string) (model.DeploymentDto, error)
	UpdateDeployment(name string, namespace string, dto model.DeploymentDto) error
	DeleteDeployment(name string, namespace string) error
	TroubleshootDeployment(name string, namespace string)
}

type deploymentUseCase struct {
	client  kubeclient.DeploymentClient
	service service.Diagnostic
}

func NewDeploymentUseCase(client kubeclient.DeploymentClient, service *service.Diagnostic) DeploymentUseCase {
	return &deploymentUseCase{client: client, service: *service}
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

func (d *deploymentUseCase) TroubleshootDeployment(name string, namespace string) {
	d.service.Analyse(name, namespace, service.Deployment)
}
