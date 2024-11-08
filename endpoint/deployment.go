package endpoint

import "github.com/beto20/kubessitant/usecase"

type DeploymentEndpoint struct {
	usecase usecase.K8sObject
}

func NewDeploymentEndpoint(usecase usecase.K8sObject) *DeploymentEndpoint {
	return &DeploymentEndpoint{}
}

func (de *DeploymentEndpoint) getDeployments() {
	de.usecase.GetAll()
}
