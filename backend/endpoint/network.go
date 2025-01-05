package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type INetworkEndpoint interface {
	GetServices(namespace string) []objects.ServiceDto
	GetIngresses(namespace string) []objects.IngressDto
	GetServiceByName(name string) objects.ServiceDto
	GetIngressByName(name string) objects.IngressDto
}

type networkEndpoint struct {
	serviceUseCase usecase.IServiceUseCase
	ingressUseCase usecase.IIngressUseCase
}

func NewNetworkEndpoint(serviceUseCase usecase.IServiceUseCase, ingressUseCase usecase.IIngressUseCase) INetworkEndpoint {
	return &networkEndpoint{serviceUseCase: serviceUseCase, ingressUseCase: ingressUseCase}
}

func (ne *networkEndpoint) GetServices(namespace string) []objects.ServiceDto {
	return ne.serviceUseCase.GetServices(namespace)
}

func (ne *networkEndpoint) GetIngresses(namespace string) []objects.IngressDto {
	return ne.ingressUseCase.GetIngresses(namespace)
}

func (ne *networkEndpoint) GetServiceByName(name string) objects.ServiceDto {
	return ne.serviceUseCase.GetServiceByName(name)
}

func (ne *networkEndpoint) GetIngressByName(name string) objects.IngressDto {
	return ne.ingressUseCase.GetIngressByName(name)
}
