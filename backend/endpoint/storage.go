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

func (se *StorageEndpoint) UpdatePersistentVolumeByName(name string, dto model.PersistentVolumeDto) error {
	return se.useCase.UpdatePersistentVolume(name, dto)
}

func (se *StorageEndpoint) DeletePersistentVolumeByName(name string) error {
	return se.useCase.DeletePersistentVolume(name)
}

func (se *StorageEndpoint) GetPersistentVolumesClaim(namespace string) ([]model.PersistentVolumeClaimDto, error) {
	return se.useCase.GetPersistentVolumesClaim(namespace)
}

func (se *StorageEndpoint) GetPersistentVolumeClaim(name string, namespace string) (model.PersistentVolumeClaimDto, error) {
	return se.useCase.GetPersistentVolumeClaim(name, namespace)
}

func (se *StorageEndpoint) UpdatePersistentVolumeClaim(name string, namespace string, dto model.PersistentVolumeClaimDto) error {
	return se.useCase.UpdatePersistentVolumeClaim(name, namespace, dto)
}

func (se *StorageEndpoint) DeletePersistentVolumeClaim(name string, namespace string) error {
	return se.useCase.DeletePersistentVolumeClaim(name, namespace)
}
