package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type INetworkEndpoint interface {
	GetServices() []objects.ServiceDto
	GetServiceByName(name string) objects.ServiceDto
	UpdateServiceByName(name string) bool
	DeleteServiceByName(name string) bool
	GetIngresses() []objects.IngressDto
	GetIngressByName(name string) objects.IngressDto
	UpdateIngressByName(name string) bool
	DeleteIngressByName(name string) bool
}

type networkEndpoint struct {
	serviceUseCase usecase.IServiceUseCase
	ingressUseCase usecase.IIngressUseCase
}

func NewNetworkEndpoint(serviceUseCase usecase.IServiceUseCase, ingressUseCase usecase.IIngressUseCase) INetworkEndpoint {
	return &networkEndpoint{serviceUseCase: serviceUseCase, ingressUseCase: ingressUseCase}
}

func (ne *networkEndpoint) GetServices() []objects.ServiceDto {
	return ne.serviceUseCase.GetServices()
}

func (ne *networkEndpoint) GetServiceByName(name string) objects.ServiceDto {
	return ne.serviceUseCase.GetServiceByName(name)
}

func (ne *networkEndpoint) UpdateServiceByName(name string) bool {
	return ne.serviceUseCase.UpdateServiceByName(name)
}

func (ne *networkEndpoint) DeleteServiceByName(name string) bool {
	return ne.serviceUseCase.DeleteServiceByName(name)
}

func (ne *networkEndpoint) GetIngresses() []objects.IngressDto {
	return ne.ingressUseCase.GetIngresses()
}

func (ne *networkEndpoint) GetIngressByName(name string) objects.IngressDto {
	return ne.ingressUseCase.GetIngressByName(name)
}

func (ne *networkEndpoint) UpdateIngressByName(name string) bool {
	return ne.ingressUseCase.UpdateIngressByName(name)
}

func (ne *networkEndpoint) DeleteIngressByName(name string) bool {
	return ne.ingressUseCase.DeleteIngressByName(name)
}
