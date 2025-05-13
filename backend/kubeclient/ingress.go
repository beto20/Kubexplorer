package kubeclient

import (
	"Kubessistant/backend/model"
	"k8s.io/client-go/kubernetes"
)

type ingressClient struct {
	client kubernetes.Interface
}

func NewIngressClient(client kubernetes.Interface) IngressClient {
	return &ingressClient{client: client}
}

func (i ingressClient) GetIngresses() ([]model.IngressDto, error) {
	//TODO implement me
	panic("implement me")
}

func (i ingressClient) GetIngress(name string) (model.IngressDto, error) {
	//TODO implement me
	panic("implement me")
}

func (i ingressClient) UpdateIngress(name string) error {
	//TODO implement me
	panic("implement me")
}

func (i ingressClient) DeleteIngress(name string) error {
	//TODO implement me
	panic("implement me")
}
