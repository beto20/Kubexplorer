package endpoint

import (
	"Kubessistant/backend/database"
	"Kubessistant/backend/usecase"
)

type IParameterEndpoint interface {
	GetKubernetesParameters() []database.CommonParameterDto
	GetCommonParameters() []database.CommonParameterDto
}

type ParameterEndpoint struct {
	useCase usecase.IParameterUseCase
}

func NewParameterEndpoint(useCase usecase.IParameterUseCase) IParameterEndpoint {
	return &ParameterEndpoint{useCase: useCase}
}

func (p *ParameterEndpoint) GetKubernetesParameters() []database.CommonParameterDto {
	return p.useCase.GetKubernetesParameters()
}

func (p *ParameterEndpoint) GetCommonParameters() []database.CommonParameterDto {
	return p.useCase.GetCommonParameters()
}
