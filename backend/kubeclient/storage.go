package kubeclient

import (
	"Kubessistant/backend/model"
	"k8s.io/client-go/kubernetes"
)

type storageClient struct {
	client kubernetes.Interface
}

func NewStorageClient(client kubernetes.Interface) StorageClient {
	return &storageClient{client: client}
}

func (s storageClient) GetPersistentVolumes() ([]model.PersistentVolumeDto, error) {
	//TODO implement me
	panic("implement me")
}

func (s storageClient) GetPersistentVolume(name string) (model.PersistentVolumeDto, error) {
	//TODO implement me
	panic("implement me")
}

func (s storageClient) UpdatePersistentVolume(name string) error {
	//TODO implement me
	panic("implement me")
}

func (s storageClient) DeletePersistentVolume(name string) error {
	//TODO implement me
	panic("implement me")
}
