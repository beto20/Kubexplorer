package usecase

import "Kubessistant/backend/objects"

type IEnvironmentUseCase interface {
	GetAllEnvironment() []objects.EnvironmentDto
	GetCurrentEnvironment(env string, name string) objects.EnvironmentDto
}

type environmentImpl struct {
	object objects.IEnvironmentObject
}

func NewEnvironmentUseCase(object objects.IEnvironmentObject) IEnvironmentUseCase {
	return &environmentImpl{object: object}
}

func (e *environmentImpl) GetCurrentEnvironment(env string, name string) objects.EnvironmentDto {
	return e.object.GetCurrentEnvironment(env, name)
}

func (e *environmentImpl) GetAllEnvironment() []objects.EnvironmentDto {
	return e.object.GetAllEnvironment()
}
