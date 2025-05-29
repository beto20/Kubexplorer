package middleware

import (
	"Kubexplorer/backend/endpoint"
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/usecase"
	"k8s.io/client-go/kubernetes"
)

type StorageMiddleware struct {
	endpoint endpoint.StorageEndpoint
}

func NewStorageMiddleware(endpoint *endpoint.StorageEndpoint) *StorageMiddleware {
	return &StorageMiddleware{endpoint: *endpoint}
}

func (s *StorageMiddleware) GetPersistentVolumes() ([]model.PersistentVolumeDto, error) {
	return s.endpoint.GetPersistentVolumes()
}

func (s *StorageMiddleware) GetPersistentVolumeByName(name string) (model.PersistentVolumeDto, error) {
	return s.endpoint.GetPersistentVolumeByName(name)
}

func (s *StorageMiddleware) UpdatePersistentVolumeByName(name string) error {
	return s.endpoint.UpdatePersistentVolumeByName(name)
}

func (s *StorageMiddleware) DeletePersistentVolumeByName(name string) error {
	return s.endpoint.DeletePersistentVolumeByName(name)
}

func BuildStorage(client kubernetes.Interface) *StorageMiddleware {
	storageClient := kubeclient.NewStorageClient(client)

	storageUseCase := usecase.NewStorageUseCase(storageClient)

	storageEndpoint := endpoint.NewStorageEndpoint(storageUseCase)

	return NewStorageMiddleware(storageEndpoint)
}
