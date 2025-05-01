package kubeclient

import (
	"Kubessistant/backend/model"
)

type ClusterClient interface {
	ListAvailableClusters(profile model.ClusterProfile) ([]model.ClusterInfo, error)
}

type MetricClient interface {
	GetPodMetrics(namespace string, chMetricDto <-chan []model.PodMetricDto) []model.PodMetricDto
}

type DeploymentClient interface {
	GetDeployments() ([]model.DeploymentDto, error)
	GetDeploymentsMock() ([]model.DeploymentDto, error)
	GetDeployment(name string) (model.DeploymentDto, error)
	UpdateDeployment(name string) error
	DeleteDeployment(name string) error
}
