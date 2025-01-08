package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type IStorageEndpoint interface {
	GetPersistentVolumes() []objects.PersistentVolumeDto
	GetPersistentVolumeByName(name string) objects.PersistentVolumeDto
	UpdatePersistentVolumeByName(name string) bool
	DeletePersistentVolumeByName(name string) bool
}

type storageEndpoint struct {
	useCase usecase.IStorageUseCase
}

func NewStorageEndpoint(useCase usecase.IStorageUseCase) IStorageEndpoint {
	return &storageEndpoint{useCase: useCase}
}

func (se *storageEndpoint) GetPersistentVolumes() []objects.PersistentVolumeDto {
	return se.useCase.GetPersistentVolumes()
}

func (se *storageEndpoint) GetPersistentVolumeByName(name string) objects.PersistentVolumeDto {
	return se.useCase.GetPersistentVolumeByName(name)
}

func (se *storageEndpoint) UpdatePersistentVolumeByName(name string) bool {
	return se.useCase.UpdatePersistentVolumeByName(name)
}

func (se *storageEndpoint) DeletePersistentVolumeByName(name string) bool {
	return se.useCase.DeletePersistentVolumeByName(name)
}
