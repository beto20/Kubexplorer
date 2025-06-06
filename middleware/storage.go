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

func (s *StorageMiddleware) UpdatePersistentVolumeByName(name string, dto model.PersistentVolumeDto) error {
	return s.endpoint.UpdatePersistentVolumeByName(name, dto)
}

func (s *StorageMiddleware) DeletePersistentVolumeByName(name string) error {
	return s.endpoint.DeletePersistentVolumeByName(name)
}

func (s *StorageMiddleware) GetPersistentVolumesClaim(namespace string) ([]model.PersistentVolumeClaimDto, error) {
	return s.endpoint.GetPersistentVolumesClaim(namespace)
}

func (s *StorageMiddleware) GetPersistentVolumeClaim(name string, namespace string) (model.PersistentVolumeClaimDto, error) {
	return s.endpoint.GetPersistentVolumeClaim(name, namespace)
}

func (s *StorageMiddleware) UpdatePersistentVolumeClaim(name string, namespace string, dto model.PersistentVolumeClaimDto) error {
	return s.endpoint.UpdatePersistentVolumeClaim(name, namespace, dto)
}

func (s *StorageMiddleware) DeletePersistentVolumeClaim(name string, namespace string) error {
	return s.endpoint.DeletePersistentVolumeClaim(name, namespace)
}

func BuildStorage(client kubernetes.Interface) *StorageMiddleware {
	storageClient := kubeclient.NewStorageClient(client)

	storageUseCase := usecase.NewStorageUseCase(storageClient)

	storageEndpoint := endpoint.NewStorageEndpoint(storageUseCase)

	return NewStorageMiddleware(storageEndpoint)
}
