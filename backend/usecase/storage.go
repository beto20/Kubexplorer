package usecase

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
)

type StorageUseCase interface {
	GetPersistentVolumes() ([]model.PersistentVolumeDto, error)
	GetPersistentVolume(name string) (model.PersistentVolumeDto, error)
	UpdatePersistentVolume(name string) error
	DeletePersistentVolume(name string) error
}

type storageUseCase struct {
	client kubeclient.StorageClient
}

func NewStorageUseCase(client kubeclient.StorageClient) StorageUseCase {
	return &storageUseCase{client: client}
}

func (s *storageUseCase) GetPersistentVolumes() ([]model.PersistentVolumeDto, error) {
	return s.client.GetPersistentVolumes()
}

func (s *storageUseCase) GetPersistentVolume(name string) (model.PersistentVolumeDto, error) {
	return s.client.GetPersistentVolume(name)
}

func (s *storageUseCase) UpdatePersistentVolume(name string) error {
	return s.client.UpdatePersistentVolume(name)
}

func (s *storageUseCase) DeletePersistentVolume(name string) error {
	return s.client.DeletePersistentVolume(name)
}
