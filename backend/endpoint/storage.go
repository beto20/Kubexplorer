package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type StorageEndpoint struct {
	useCase usecase.IStorageUseCase
}

func NewStorageEndpoint(useCase usecase.IStorageUseCase) *StorageEndpoint {
	return &StorageEndpoint{useCase: useCase}
}

func (se *StorageEndpoint) GetPersistentVolumes() []objects.PersistentVolumeDto {
	return se.useCase.GetPersistentVolumes()
}

func (se *StorageEndpoint) GetPersistentVolumeByName(name string) objects.PersistentVolumeDto {
	return se.useCase.GetPersistentVolumeByName(name)
}

func (se *StorageEndpoint) UpdatePersistentVolumeByName(name string) bool {
	return se.useCase.UpdatePersistentVolumeByName(name)
}

func (se *StorageEndpoint) DeletePersistentVolumeByName(name string) bool {
	return se.useCase.DeletePersistentVolumeByName(name)
}
