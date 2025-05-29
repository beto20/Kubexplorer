package usecase

import (
	"Kubexplorer/backend/service"
)

type ResourceUseCase interface {
	Invoke(namespace string)
}

type resourceUseCase struct {
	service service.ResourceService
}

func NewResourceUseCase(service service.ResourceService) ResourceUseCase {
	return &resourceUseCase{service: service}
}

func (r *resourceUseCase) Invoke(namespace string) {
	r.service.ResourceTuning(namespace)
}
