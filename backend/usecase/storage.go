package usecase

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
)

type StorageUseCase interface {
	GetPersistentVolumes() ([]model.PersistentVolumeDto, error)
	GetPersistentVolume(name string) (model.PersistentVolumeDto, error)
	UpdatePersistentVolume(name string, dto model.PersistentVolumeDto) error
	DeletePersistentVolume(name string) error

	GetPersistentVolumesClaim(namespace string) ([]model.PersistentVolumeClaimDto, error)
	GetPersistentVolumeClaim(name string, namespace string) (model.PersistentVolumeClaimDto, error)
	UpdatePersistentVolumeClaim(name string, namespace string, dto model.PersistentVolumeClaimDto) error
	DeletePersistentVolumeClaim(name string, namespace string) error
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

func (s *storageUseCase) UpdatePersistentVolume(name string, dto model.PersistentVolumeDto) error {
	return s.client.UpdatePersistentVolume(name, dto)
}

func (s *storageUseCase) DeletePersistentVolume(name string) error {
	return s.client.DeletePersistentVolume(name)
}

func (s *storageUseCase) GetPersistentVolumesClaim(namespace string) ([]model.PersistentVolumeClaimDto, error) {
	return s.client.GetPersistentVolumesClaim(namespace)
}

func (s *storageUseCase) GetPersistentVolumeClaim(name string, namespace string) (model.PersistentVolumeClaimDto, error) {
	return s.client.GetPersistentVolumeClaim(name, namespace)
}

func (s *storageUseCase) UpdatePersistentVolumeClaim(name string, namespace string, dto model.PersistentVolumeClaimDto) error {
	return s.client.UpdatePersistentVolumeClaim(name, namespace, dto)
}

func (s *storageUseCase) DeletePersistentVolumeClaim(name string, namespace string) error {
	return s.client.DeletePersistentVolumeClaim(name, namespace)
}
