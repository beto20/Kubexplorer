package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type IGeneralEndpoint interface {
	GetNodes() []objects.NodeDto
	GetNodeByName(name string) objects.NodeDto
	GetNamespace() []objects.NamespaceDto
	GetNamespaceByName(name string) objects.NamespaceDto
	UpdateNamespaceByName(name string) bool
	DeleteNamespaceByName(name string) bool
}

type generalEndpoint struct {
	nodeUseCase      usecase.INodeUseCase
	namespaceUseCase usecase.INamespaceUseCase
}

func NewGeneralEndpoint(nodeUseCase usecase.INodeUseCase, namespaceUseCase usecase.INamespaceUseCase) IGeneralEndpoint {
	return &generalEndpoint{
		nodeUseCase:      nodeUseCase,
		namespaceUseCase: namespaceUseCase,
	}
}

func (ge *generalEndpoint) GetNodes() []objects.NodeDto {
	return ge.nodeUseCase.GetNodes()
}

func (ge *generalEndpoint) GetNodeByName(name string) objects.NodeDto {
	return ge.nodeUseCase.GetNodeByName(name)
}

func (ge *generalEndpoint) GetNamespace() []objects.NamespaceDto {
	return ge.namespaceUseCase.GetNamespace()
}

func (ge *generalEndpoint) GetNamespaceByName(name string) objects.NamespaceDto {
	return ge.namespaceUseCase.GetNamespaceByName(name)
}

func (ge *generalEndpoint) UpdateNamespaceByName(name string) bool {
	return ge.namespaceUseCase.UpdateNamespaceByName(name)
}

func (ge *generalEndpoint) DeleteNamespaceByName(name string) bool {
	return ge.namespaceUseCase.DeleteNamespaceByName(name)
}
