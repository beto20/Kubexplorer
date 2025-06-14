package endpoint

import (
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/usecase"
	"fmt"
)

type EnvironmentEndpoint struct {
	useCase usecase.EnvironmentUseCase
}

func NewEnvironmentEndpoint(useCase usecase.EnvironmentUseCase) *EnvironmentEndpoint {
	return &EnvironmentEndpoint{useCase: useCase}
}

func (ee *EnvironmentEndpoint) GetCurrentEnvironment(env string, name string) (model.EnvironmentDto, error) {
	return ee.useCase.GetCurrentEnvironment(env, name)
}

func (ee *EnvironmentEndpoint) GetAllEnvironment() ([]model.EnvironmentDto, error) {
	x, _ := ee.useCase.GetAllEnvironment()
	fmt.Print("ENV NAME: ", x[2].Name)
	return ee.useCase.GetAllEnvironment()
}

func (ee *EnvironmentEndpoint) GetObjectsView() (model.ObjectMapDto, error) {
	return ee.useCase.GetObjectsView()
}
