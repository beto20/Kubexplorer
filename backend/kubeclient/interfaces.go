package kubeclient

import (
	"Kubessistant/backend/model"
)

type ClusterClient interface {
	ListAvailableClusters(profile model.ClusterProfile) ([]model.ClusterInfo, error)
}

type MetricClient interface {
	GetPodMetrics(namespace string) []model.PodMetricDto
}
