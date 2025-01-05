package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type IStorageEndpoint interface {
	GetPersistentVolumes() []objects.PersistentVolumeDto
	GetPersistentVolumeByName(name string) objects.PersistentVolumeDto
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
