package usecase

import "Kubessistant/backend/objects"

type INamespaceUseCase interface {
	GetNamespace() []objects.NamespaceDto
	GetNamespaceByName(name string) objects.NamespaceDto
	UpdateNamespaceByName(name string) bool
	DeleteNamespaceByName(name string) bool
}

type namespaceImpl struct {
	object objects.INamespaceObject
}

func NewNamespaceUseCase(object objects.INamespaceObject) INamespaceUseCase {
	return &namespaceImpl{object: object}
}

func (n *namespaceImpl) GetNamespace() []objects.NamespaceDto {
	return n.object.GetNamespace()
}
func (n *namespaceImpl) GetNamespaceByName(name string) objects.NamespaceDto {
	return n.object.GetNamespaceByName(name)
}

func (n *namespaceImpl) UpdateNamespaceByName(name string) bool {
	return n.object.UpdateNamespaceByName(name)
}

func (n *namespaceImpl) DeleteNamespaceByName(name string) bool {
	return n.object.DeleteNamespaceByName(name)
}
