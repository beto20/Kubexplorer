package usecase

import "Kubessistant/backend/objects"

type IDeploymentUseCase interface {
	GetAllDeployments(namespace string) []objects.DeploymentDto
	GetDeploymentByName(name string) objects.DeploymentDto
}

type deploymentImpl struct {
	object objects.IDeploymentObject
}

func NewDeploymentUseCase(object objects.IDeploymentObject) IDeploymentUseCase {
	return &deploymentImpl{object: object}
}

func (d *deploymentImpl) GetAllDeployments(namespace string) []objects.DeploymentDto {
	return d.object.GetDeploymentsMock(namespace)
}

func (d *deploymentImpl) GetDeploymentByName(name string) objects.DeploymentDto {
	return d.object.GetDeploymentByName(name)
}
