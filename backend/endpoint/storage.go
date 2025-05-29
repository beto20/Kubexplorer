package endpoint

import (
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/usecase"
)

type StorageEndpoint struct {
	useCase usecase.StorageUseCase
}

func NewStorageEndpoint(useCase usecase.StorageUseCase) *StorageEndpoint {
	return &StorageEndpoint{useCase: useCase}
}

func (se *StorageEndpoint) GetPersistentVolumes() ([]model.PersistentVolumeDto, error) {
	return se.useCase.GetPersistentVolumes()
}

func (se *StorageEndpoint) GetPersistentVolumeByName(name string) (model.PersistentVolumeDto, error) {
	return se.useCase.GetPersistentVolume(name)
}

func (se *StorageEndpoint) UpdatePersistentVolumeByName(name string) error {
	return se.useCase.UpdatePersistentVolume(name)
}

func (se *StorageEndpoint) DeletePersistentVolumeByName(name string) error {
	return se.useCase.DeletePersistentVolume(name)
}
