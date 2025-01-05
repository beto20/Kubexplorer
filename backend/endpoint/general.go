package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type IGeneralEndpoint interface {
	GetNodes() []objects.NodeDto
}

type generalEndpoint struct {
	useCase usecase.INodeUseCase
}

func NewGeneralEndpoint(useCase usecase.INodeUseCase) IGeneralEndpoint {
	return &generalEndpoint{useCase: useCase}
}

func (n *generalEndpoint) GetNodes() []objects.NodeDto {
	return n.useCase.GetNodes()
}
