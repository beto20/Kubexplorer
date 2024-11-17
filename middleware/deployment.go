package middleware

import (
	"Kubessistant/backend/endpoint"
	"Kubessistant/backend/objects"
)

type DeploymentMiddleware struct {
	endpoint endpoint.IDeploymentEndpoint
}

func NewDeploymentMiddleware(endpoint endpoint.IDeploymentEndpoint) *DeploymentMiddleware {
	return &DeploymentMiddleware{endpoint: endpoint}
}

func (d *DeploymentMiddleware) GetDeployments(namespace string) []objects.DeploymentDto {
	return d.endpoint.GetDeployments(namespace)
}
