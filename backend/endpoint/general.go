package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type IGeneralEndpoint interface {
	GetNodes() []objects.NodeDto
	GetNodeByName(name string) objects.NodeDto
}

type generalEndpoint struct {
	useCase usecase.INodeUseCase
}

func NewGeneralEndpoint(useCase usecase.INodeUseCase) IGeneralEndpoint {
	return &generalEndpoint{useCase: useCase}
}

func (ge *generalEndpoint) GetNodes() []objects.NodeDto {
	return ge.useCase.GetNodes()
}

func (ge *generalEndpoint) GetNodeByName(name string) objects.NodeDto {
	return ge.useCase.GetNodeByName(name)
}
