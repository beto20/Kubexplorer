package endpoint

import (
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/usecase"
)

type NetworkEndpoint struct {
	serviceUseCase usecase.ServiceUseCase
	ingressUseCase usecase.IngressUseCase
}

func NewNetworkEndpoint(serviceUseCase usecase.ServiceUseCase, ingressUseCase usecase.IngressUseCase) *NetworkEndpoint {
	return &NetworkEndpoint{serviceUseCase: serviceUseCase, ingressUseCase: ingressUseCase}
}

func (ne *NetworkEndpoint) GetServices(namespace string) ([]model.ServiceDto, error) {
	return ne.serviceUseCase.GetServices(namespace)
}

func (ne *NetworkEndpoint) GetServiceByName(name string, namespace string) (model.ServiceDto, error) {
	return ne.serviceUseCase.GetService(name, namespace)
}

func (ne *NetworkEndpoint) UpdateServiceByName(name string, namespace string, dto model.ServiceDto) error {
	return ne.serviceUseCase.UpdateService(name, namespace, dto)
}

func (ne *NetworkEndpoint) DeleteServiceByName(name string, namespace string) error {
	return ne.serviceUseCase.DeleteService(name, namespace)
}

func (ne *NetworkEndpoint) GetIngresses(namespace string) ([]model.IngressDto, error) {
	return ne.ingressUseCase.GetIngresses(namespace)
}

func (ne *NetworkEndpoint) GetIngressByName(name string, namespace string) (model.IngressDto, error) {
	return ne.ingressUseCase.GetIngress(name, namespace)
}

func (ne *NetworkEndpoint) UpdateIngressByName(name string, namespace string, dto model.IngressDto) error {
	return ne.ingressUseCase.UpdateIngress(name, namespace, dto)
}

func (ne *NetworkEndpoint) DeleteIngressByName(name string, namespace string) error {
	return ne.ingressUseCase.DeleteIngress(name, namespace)
}
