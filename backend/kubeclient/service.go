package kubeclient

import (
	"Kubessistant/backend/model"
	"k8s.io/client-go/kubernetes"
)

type serviceClient struct {
	client kubernetes.Interface
}

func NewServiceClient(client kubernetes.Interface) ServiceClient {
	return &serviceClient{client: client}
}

func (s serviceClient) GetServices(namespace string) ([]model.ServiceDto, error) {
	//TODO implement me
	panic("implement me")
}

func (s serviceClient) GetServicesMock() ([]model.ServiceDto, error) {
	//TODO implement me
	panic("implement me")
}

func (s serviceClient) GetService(name string) (model.ServiceDto, error) {
	//TODO implement me
	panic("implement me")
}

func (s serviceClient) UpdateService(name string) error {
	//TODO implement me
	panic("implement me")
}

func (s serviceClient) DeleteService(name string) error {
	//TODO implement me
	panic("implement me")
}
