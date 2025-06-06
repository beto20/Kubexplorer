package service

import (
	"Kubexplorer/backend/model"
	"k8s.io/client-go/dynamic"
)

type ClusterService interface {
	GetObjectsView() (model.ObjectMapDto, error)
}

type clusterService struct {
	client dynamic.Interface
}

func NewClusterService(client dynamic.Interface) ClusterService {
	return &clusterService{client: client}
}

func (c *clusterService) GetObjectsView() (model.ObjectMapDto, error) {
	return model.ObjectMapDto{}, nil
}
