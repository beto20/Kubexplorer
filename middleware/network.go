package middleware

import (
	"Kubessistant/backend/endpoint"
	"Kubessistant/backend/kubeclient"
	"Kubessistant/backend/model"
	"Kubessistant/backend/usecase"
	"k8s.io/client-go/kubernetes"
)

type NetworkMiddleware struct {
	endpoint endpoint.NetworkEndpoint
}

func NewNetworkMiddleware(endpoint *endpoint.NetworkEndpoint) *NetworkMiddleware {
	return &NetworkMiddleware{endpoint: *endpoint}
}

func (n *NetworkMiddleware) GetServices() ([]model.ServiceDto, error) {
	return n.endpoint.GetServices()
}

func (n *NetworkMiddleware) GetServiceByName(name string) (model.ServiceDto, error) {
	return n.endpoint.GetServiceByName(name)
}

func (n *NetworkMiddleware) UpdateServiceByName(name string) error {
	return n.endpoint.UpdateServiceByName(name)
}

func (n *NetworkMiddleware) DeleteServiceByName(name string) error {
	return n.endpoint.DeleteServiceByName(name)
}

func (n *NetworkMiddleware) GetIngresses() ([]model.IngressDto, error) {
	return n.endpoint.GetIngresses()
}

func (n *NetworkMiddleware) GetIngressByName(name string) (model.IngressDto, error) {
	return n.endpoint.GetIngressByName(name)
}

func (n *NetworkMiddleware) UpdateIngressByName(name string) error {
	return n.endpoint.UpdateIngressByName(name)
}

func (n *NetworkMiddleware) DeleteIngressByName(name string) error {
	return n.endpoint.DeleteIngressByName(name)
}

func BuildNetwork(client kubernetes.Interface) *NetworkMiddleware {
	serviceClient := kubeclient.NewServiceClient(client)
	ingressClient := kubeclient.NewIngressClient(client)

	serviceUseCase := usecase.NewServiceUseCase(serviceClient)
	ingressUseCase := usecase.NewIngressUseCase(ingressClient)

	networkEndpoint := endpoint.NewNetworkEndpoint(serviceUseCase, ingressUseCase)

	return NewNetworkMiddleware(networkEndpoint)
}
