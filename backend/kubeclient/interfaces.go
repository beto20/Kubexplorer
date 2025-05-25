package kubeclient

import (
	"Kubessistant/backend/model"
)

type ClusterClient interface {
	ListAvailableClusters(profile model.ClusterProfile) ([]model.ClusterInfo, error)
	GetCurrentCluster() (model.EnvironmentDto, error)
	GetClusters() ([]model.EnvironmentDto, error)
	GetNode(name string) (model.NodeDtoV2, error)
	GetNodes() ([]model.NodeDto, error)
}

type MetricClient interface {
	GetPodMetrics(namespace string, chMetricDto <-chan []model.PodMetricDto) []model.PodMetricDto
}

type PodClient interface {
	GetPodsMock() ([]model.PodDto, error)
	GetPods() ([]model.PodDto, error)
	GetPod(name string, namespace string) (model.PodDto, error)
	UpdatePod(name string, namespace string, dto model.PodDto) error
	DeletePod(name string, namespace string) error
}

type DeploymentClient interface {
	GetDeployments() ([]model.DeploymentDto, error)
	GetDeploymentsMock() ([]model.DeploymentDto, error)
	GetDeployment(name string, namespace string) (model.DeploymentDto, error)
	UpdateDeployment(name string, namespace string, dto model.DeploymentDto) error
	DeleteDeployment(name string, namespace string) error
}

type StorageClient interface {
	GetPersistentVolumes() ([]model.PersistentVolumeDto, error)
	GetPersistentVolume(name string) (model.PersistentVolumeDto, error)
	UpdatePersistentVolume(name string) error
	DeletePersistentVolume(name string) error
}

type NamespaceClient interface {
	GetNamespaces() ([]model.NamespaceDto, error)
	GetNamespace(name string) (model.NamespaceDto, error)
	UpdateNamespace(name string, dto model.NamespaceDto) error
	DeleteNamespace(name string) error
}

type ServiceClient interface {
	GetServices(namespace string) ([]model.ServiceDto, error)
	GetService(name string, namespace string) (model.ServiceDto, error)
	UpdateService(name string, namespace string, dto model.ServiceDto) error
	DeleteService(name string, namespace string) error
}

type IngressClient interface {
	GetIngresses() ([]model.IngressDto, error)
	GetIngress(name string, namespace string) (model.IngressDto, error)
	UpdateIngress(name string, namespace string, dto model.IngressDto) error
	DeleteIngress(name string, namespace string) error
}
