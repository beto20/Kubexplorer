package usecase

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
)

type EnvironmentUseCase interface {
	GetAllEnvironment() ([]model.EnvironmentDto, error)
	GetCurrentEnvironment(env string, name string) (model.EnvironmentDto, error)
}

type environmentUseCase struct {
	client kubeclient.ClusterClient
}

func NewEnvironmentUseCase(client kubeclient.ClusterClient) EnvironmentUseCase {
	return &environmentUseCase{client: client}
}

func (e *environmentUseCase) GetCurrentEnvironment(env string, name string) (model.EnvironmentDto, error) {
	return e.client.GetCurrentCluster()
}

func (e *environmentUseCase) GetAllEnvironment() ([]model.EnvironmentDto, error) {
	return e.client.GetClusters()
}
