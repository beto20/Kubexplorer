package endpoint

import (
	"Kubessistant/backend/database"
	"Kubessistant/backend/usecase"
)

type IParameterEndpoint interface {
	GetKubernetesParameters() []database.CommonParameterDto
	GetCommonParameters() []database.CommonParameterDto
}

type parameterEndpoint struct {
	useCase usecase.IParameterUseCase
}

func NewParameterEndpoint(useCase usecase.IParameterUseCase) IParameterEndpoint {
	return &parameterEndpoint{useCase: useCase}
}

func (p *parameterEndpoint) GetKubernetesParameters() []database.CommonParameterDto {
	return p.useCase.GetKubernetesParameters()
}

func (p *parameterEndpoint) GetCommonParameters() []database.CommonParameterDto {
	return p.useCase.GetCommonParameters()
}
