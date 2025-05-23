package endpoint

import (
	"Kubessistant/backend/model"
	"Kubessistant/backend/usecase"
)

type GeneralEndpoint struct {
	nodeUseCase      usecase.NodeUseCase
	namespaceUseCase usecase.NamespaceUseCase
}

func NewGeneralEndpoint(nodeUseCase usecase.NodeUseCase, namespaceUseCase usecase.NamespaceUseCase) *GeneralEndpoint {
	return &GeneralEndpoint{
		nodeUseCase:      nodeUseCase,
		namespaceUseCase: namespaceUseCase,
	}
}

func (ge *GeneralEndpoint) GetNodes() ([]model.NodeDto, error) {
	return ge.nodeUseCase.GetNodes()
}

func (ge *GeneralEndpoint) GetNodeByName(name string) (model.NodeDto, error) {
	return ge.nodeUseCase.GetNode(name)
}

func (ge *GeneralEndpoint) GetNamespace() ([]model.NamespaceDto, error) {
	return ge.namespaceUseCase.GetNamespaces()
}

func (ge *GeneralEndpoint) GetNamespaceByName(name string) (model.NamespaceDto, error) {
	return ge.namespaceUseCase.GetNamespace(name)
}

func (ge *GeneralEndpoint) UpdateNamespaceByName(name string) error {
	return ge.namespaceUseCase.UpdateNamespace(name)
}

func (ge *GeneralEndpoint) DeleteNamespaceByName(name string) error {
	return ge.namespaceUseCase.DeleteNamespace(name)
}
