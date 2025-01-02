package usecase

import "Kubessistant/backend/database"

type IParameterUseCase interface {
	GetKubernetesParameters() []database.CommonParameterDto
	GetCommonParameters() []database.CommonParameterDto
	GetK8sObjects() []database.ObjectType
}

type parameterImpl struct {
	entity database.IParameterEntity
}

func NewParameterUseCase(entity database.IParameterEntity) IParameterUseCase {
	return &parameterImpl{entity: entity}
}

func (p *parameterImpl) GetKubernetesParameters() []database.CommonParameterDto {
	return p.entity.GetKubernetesParameters()
}

func (p *parameterImpl) GetCommonParameters() []database.CommonParameterDto {
	return p.entity.GetCommonParameters()
}

func (p *parameterImpl) GetK8sObjects() []database.ObjectType {
	return p.entity.GetK8sObjects()
}
