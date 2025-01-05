package middleware

import (
	"Kubessistant/backend/endpoint"
	"Kubessistant/backend/objects"
)

type StorageMiddleware struct {
	endpoint endpoint.IStorageEndpoint
}

func NewStorageMiddleware(endpoint endpoint.IStorageEndpoint) *StorageMiddleware {
	return &StorageMiddleware{endpoint: endpoint}
}

func (s *StorageMiddleware) GetPersistentVolumes() []objects.PersistentVolumeDto {
	return s.endpoint.GetPersistentVolumes()
}

func (s *StorageMiddleware) GetPersistentVolumeByName(name string) objects.PersistentVolumeDto {
	return s.endpoint.GetPersistentVolumeByName(name)
}
