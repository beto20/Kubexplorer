package kubeclient

import (
	"Kubessistant/backend/model"
)

type PodClient interface {
	GetPodsMock() ([]model.PodDto, error)
	GetPods() ([]model.PodDto, error)
	GetPod(name string) (model.PodDto, error)
	UpdatePod(name string) error
	DeletePod(name string) error
	GetPodMetric(namespace string) []model.PodMetricDto
}

type DeploymentClient interface {
	GetDeployments() ([]model.DeploymentDto, error)
	GetDeploymentsMock() ([]model.DeploymentDto, error)
	GetDeployment(name string) (model.DeploymentDto, error)
	UpdateDeployment(name string) error
	DeleteDeployment(name string) error
}
