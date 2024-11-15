package usecase

import "github.com/beto20/kubessitant/objects"

type IDeploymentUseCase interface {
	GetAllDeployments(namespace string) []objects.DeploymentDto
}

type deploymentImpl struct {
	object objects.IDeployment
}

func NewDeploymentUseCase(object objects.IDeployment) IDeploymentUseCase {
	return &deploymentImpl{object: object}
}

func (d *deploymentImpl) GetAllDeployments(namespace string) []objects.DeploymentDto {
	return d.object.GetDeploymentsMock(namespace)
}
