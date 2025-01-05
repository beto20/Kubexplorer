package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type IStorageEndpoint interface {
	GetPersistentVolumes() []objects.PersistentVolumeDto
}

type storageEndpoint struct {
	useCase usecase.IStorageUseCase
}

func NewStorageEndpoint(useCase usecase.IStorageUseCase) IStorageEndpoint {
	return &storageEndpoint{useCase: useCase}
}

func (s *storageEndpoint) GetPersistentVolumes() []objects.PersistentVolumeDto {
	return s.useCase.GetPersistentVolumes()
}
