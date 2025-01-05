package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type INodeEndpoint interface {
	GetNodes() []objects.NodeDto
}

type NodeEndpoint struct {
	useCase usecase.INodeUseCase
}

func NewNodeEndpoint(useCase usecase.INodeUseCase) INodeEndpoint {
	return &NodeEndpoint{useCase}
}

func (n *NodeEndpoint) GetNodes() []objects.NodeDto {
	return n.useCase.GetNodes()
}
