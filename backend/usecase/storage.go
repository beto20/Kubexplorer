package usecase

import "Kubessistant/backend/objects"

type IStorageUseCase interface {
	GetPersistentVolumes() []objects.PersistentVolumeDto
	GetPersistentVolumeByName(name string) objects.PersistentVolumeDto
	UpdatePersistentVolumeByName(name string) bool
	DeletePersistentVolumeByName(name string) bool
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

func (s *storageImpl) GetPersistentVolumeByName(name string) objects.PersistentVolumeDto {
	return s.object.GetPersistentVolumeByName(name)
}

func (s *storageImpl) UpdatePersistentVolumeByName(name string) bool {
	return s.object.UpdatePersistentVolumeByName(name)
}

func (s *storageImpl) DeletePersistentVolumeByName(name string) bool {
	return s.object.DeletePersistentVolumeByName(name)
}
