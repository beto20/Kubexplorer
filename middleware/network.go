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

func (n *NetworkMiddleware) GetServices(namespace string) ([]model.ServiceDto, error) {
	return n.endpoint.GetServices(namespace)
}

func (n *NetworkMiddleware) GetServiceByName(name string, namespace string) (model.ServiceDto, error) {
	return n.endpoint.GetServiceByName(name, namespace)
}

func (n *NetworkMiddleware) UpdateServiceByName(name string, namespace string, dto model.ServiceDto) error {
	return n.endpoint.UpdateServiceByName(name, namespace, dto)
}

func (n *NetworkMiddleware) DeleteServiceByName(name string, namespace string) error {
	return n.endpoint.DeleteServiceByName(name, namespace)
}

func (n *NetworkMiddleware) GetIngresses(namespace string) ([]model.IngressDto, error) {
	return n.endpoint.GetIngresses(namespace)
}

func (n *NetworkMiddleware) GetIngressByName(name string, namespace string) (model.IngressDto, error) {
	return n.endpoint.GetIngressByName(name, namespace)
}

func (n *NetworkMiddleware) UpdateIngressByName(name string, namespace string, dto model.IngressDto) error {
	return n.endpoint.UpdateIngressByName(name, namespace, dto)
}

func (n *NetworkMiddleware) DeleteIngressByName(name string, namespace string) error {
	return n.endpoint.DeleteIngressByName(name, namespace)
}

func BuildNetwork(client kubernetes.Interface) *NetworkMiddleware {
	serviceClient := kubeclient.NewServiceClient(client)
	ingressClient := kubeclient.NewIngressClient(client)

	serviceUseCase := usecase.NewServiceUseCase(serviceClient)
	ingressUseCase := usecase.NewIngressUseCase(ingressClient)

	networkEndpoint := endpoint.NewNetworkEndpoint(serviceUseCase, ingressUseCase)

	return NewNetworkMiddleware(networkEndpoint)
}
