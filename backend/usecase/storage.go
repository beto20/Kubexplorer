package usecase

import "Kubessistant/backend/objects"

type IStorageUseCase interface {
	GetPersistentVolumes() []objects.PersistentVolumeDto
}

type storageImpl struct {
	object objects.IStorageObject
}

func NewStorageUseCase(object objects.IStorageObject) IStorageUseCase {
	return &storageImpl{object: object}
}

func (s *storageImpl) GetPersistentVolumes() []objects.PersistentVolumeDto {
	return s.object.GetPersistentVolumes()
}
