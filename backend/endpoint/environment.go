package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
	"fmt"
)

type IEnvironmentEndpoint interface {
	GetCurrentEnvironment(env string, name string) objects.EnvironmentDto
	GetAllEnvironment() []objects.EnvironmentDto
}

type environmentEndpoint struct {
	useCase usecase.IEnvironmentUseCase
}

func NewEnvironmentEndpoint(useCase usecase.IEnvironmentUseCase) IEnvironmentEndpoint {
	return &environmentEndpoint{useCase: useCase}
}

func (ee *environmentEndpoint) GetCurrentEnvironment(env string, name string) objects.EnvironmentDto {
	return ee.useCase.GetCurrentEnvironment(env, name)
}

func (ee *environmentEndpoint) GetAllEnvironment() []objects.EnvironmentDto {
	x := ee.useCase.GetAllEnvironment()
	fmt.Print("ENV NAME: ", x[2].Name)
	return ee.useCase.GetAllEnvironment()
}
