package usecase

import (
	"Kubessistant/backend/kubeclient"
	"Kubessistant/backend/model"
)

type NodeUseCase interface {
	GetNodes() ([]model.NodeDto, error)
	GetNode(name string) (model.NodeDtoV2, error)
}

type nodeUseCase struct {
	client kubeclient.ClusterClient
}

func NewNodeUseCase(client kubeclient.ClusterClient) NodeUseCase {
	return &nodeUseCase{client: client}
}

func (n *nodeUseCase) GetNodes() ([]model.NodeDto, error) {
	return n.client.GetNodes()
}

func (n *nodeUseCase) GetNode(name string) (model.NodeDtoV2, error) {
	return n.client.GetNode(name)
}
