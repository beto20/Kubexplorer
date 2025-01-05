package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type INodeEndpoint interface {
	GetNodes() []objects.NodeDto
}

type nodeEndpoint struct {
	useCase usecase.INodeUseCase
}

func NewNodeEndpoint(useCase usecase.INodeUseCase) INodeEndpoint {
	return &nodeEndpoint{useCase}
}

func (n *nodeEndpoint) GetNodes() []objects.NodeDto {
	return n.useCase.GetNodes()
}
