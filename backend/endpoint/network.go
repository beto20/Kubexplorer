package endpoint

import (
	"Kubessistant/backend/model"
	"Kubessistant/backend/usecase"
)

type NetworkEndpoint struct {
	serviceUseCase usecase.ServiceUseCase
	ingressUseCase usecase.IngressUseCase
}

func NewNetworkEndpoint(serviceUseCase usecase.ServiceUseCase, ingressUseCase usecase.IngressUseCase) *NetworkEndpoint {
	return &NetworkEndpoint{serviceUseCase: serviceUseCase, ingressUseCase: ingressUseCase}
}

func (ne *NetworkEndpoint) GetServices() ([]model.ServiceDto, error) {
	return ne.serviceUseCase.GetServices()
}

func (ne *NetworkEndpoint) GetServiceByName(name string) (model.ServiceDto, error) {
	return ne.serviceUseCase.GetService(name)
}

func (ne *NetworkEndpoint) UpdateServiceByName(name string) error {
	return ne.serviceUseCase.UpdateService(name)
}

func (ne *NetworkEndpoint) DeleteServiceByName(name string) error {
	return ne.serviceUseCase.DeleteService(name)
}

func (ne *NetworkEndpoint) GetIngresses() ([]model.IngressDto, error) {
	return ne.ingressUseCase.GetIngresses()
}

func (ne *NetworkEndpoint) GetIngressByName(name string) (model.IngressDto, error) {
	return ne.ingressUseCase.GetIngress(name)
}

func (ne *NetworkEndpoint) UpdateIngressByName(name string) error {
	return ne.ingressUseCase.UpdateIngress(name)
}

func (ne *NetworkEndpoint) DeleteIngressByName(name string) error {
	return ne.ingressUseCase.DeleteIngress(name)
}
