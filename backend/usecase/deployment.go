package usecase

import "Kubessistant/backend/objects"

type IDeploymentUseCase interface {
	GetAllDeployments() []objects.DeploymentDto
	GetDeploymentByName(name string) objects.DeploymentDto
	UpdateDeploymentByName(name string) bool
	DeleteDeploymentByName(name string) bool
}

type deploymentImpl struct {
	object objects.IDeploymentObject
}

func NewDeploymentUseCase(object objects.IDeploymentObject) IDeploymentUseCase {
	return &deploymentImpl{object: object}
}

func (d *deploymentImpl) GetAllDeployments() []objects.DeploymentDto {
	return d.object.GetDeploymentsMock()
}

func (d *deploymentImpl) GetDeploymentByName(name string) objects.DeploymentDto {
	return d.object.GetDeploymentByName(name)
}

func (d *deploymentImpl) UpdateDeploymentByName(name string) bool {
	return d.object.UpdateDeploymentByName(name)
}

func (d *deploymentImpl) DeleteDeploymentByName(name string) bool {
	return d.object.DeleteDeploymentByName(name)
}
