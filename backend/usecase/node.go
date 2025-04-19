package usecase

import "Kubessistant/backend/objects"

type INodeUseCase interface {
	GetNodes() []objects.NodeDto
	GetNodeByName(name string) objects.NodeDto
}

type nodeImpl struct {
	object objects.INodeObject
}

func NewNodeUseCase(object objects.INodeObject) INodeUseCase {
	return &nodeImpl{object: object}
}

func (n *nodeImpl) GetNodes() []objects.NodeDto {
	return n.object.GetNodes()
}

func (n *nodeImpl) GetNodeByName(name string) objects.NodeDto {
	return n.object.GetNodeByName(name)
}
