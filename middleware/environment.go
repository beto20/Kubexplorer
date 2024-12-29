package middleware

import (
	"Kubessistant/backend/endpoint"
	"Kubessistant/backend/objects"
)

type EnvironmentMiddleware struct {
	endpoint endpoint.IEnvironmentEndpoint
}

func NewEnvironmentMiddleware(endpoint endpoint.IEnvironmentEndpoint) *EnvironmentMiddleware {
	return &EnvironmentMiddleware{endpoint}
}

func (e *EnvironmentMiddleware) GetCurrentEnvironment(env string, name string) objects.EnvironmentDto {
	return e.endpoint.GetCurrentEnvironment(env, name)
}

func (e *EnvironmentMiddleware) GetAllEnvironment() []objects.EnvironmentDto {
	return e.endpoint.GetAllEnvironment()
}
