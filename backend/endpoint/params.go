package endpoint

import (
	"Kubessistant/backend/database"
	"Kubessistant/backend/usecase"
)

type IParameterEndpoint interface {
	GetKubernetesParameters() []database.CommonParameterDto
	GetCommonParameters() []database.CommonParameterDto
	GetK8sObjects() []database.ObjectType
}

type parameterEndpoint struct {
	useCase usecase.IParameterUseCase
}

func NewParameterEndpoint(useCase usecase.IParameterUseCase) IParameterEndpoint {
	return &parameterEndpoint{useCase: useCase}
}

func (pe *parameterEndpoint) GetKubernetesParameters() []database.CommonParameterDto {
	return pe.useCase.GetKubernetesParameters()
}

func (pe *parameterEndpoint) GetCommonParameters() []database.CommonParameterDto {
	return pe.useCase.GetCommonParameters()
}

func (pe *parameterEndpoint) GetK8sObjects() []database.ObjectType {
	return pe.useCase.GetK8sObjects()
}
