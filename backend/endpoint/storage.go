package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type IStorageEndpoint interface {
	GetPersistentVolumes() []objects.PersistentVolumeDto
}

type StorageEndpoint struct {
	useCase usecase.IStorageUseCase
}

func NewStorageEndpoint(useCase usecase.IStorageUseCase) IStorageEndpoint {
	return &StorageEndpoint{useCase: useCase}
}

func (s *StorageEndpoint) GetPersistentVolumes() []objects.PersistentVolumeDto {
	return s.useCase.GetPersistentVolumes()
}
