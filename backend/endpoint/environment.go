package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
	"fmt"
)

type IEnvironmentEndpoint interface {
	GetCurrentEnvironment(envName string) objects.EnvironmentDto
	GetAllEnvironment() []objects.EnvironmentDto
}

type EnvironmentEndpoint struct {
	useCase usecase.IEnvironmentUseCase
}

func NewEnvironmentEndpoint(useCase usecase.IEnvironmentUseCase) IEnvironmentEndpoint {
	return &EnvironmentEndpoint{useCase: useCase}
}

func (ee *EnvironmentEndpoint) GetCurrentEnvironment(envName string) objects.EnvironmentDto {
	return ee.useCase.GetCurrentEnvironment(envName)
}

func (ee *EnvironmentEndpoint) GetAllEnvironment() []objects.EnvironmentDto {
	x := ee.useCase.GetAllEnvironment()
	fmt.Print("ENV NAME:", x[0].Name)
	return ee.useCase.GetAllEnvironment()
}
