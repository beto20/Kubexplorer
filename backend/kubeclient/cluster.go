package kubeclient

import (
	"Kubessistant/backend/model"
	"k8s.io/client-go/tools/clientcmd"
)

type cluster struct{}

func NewCluster() ClusterClient {
	return &cluster{}
}

func (c *cluster) ListAvailableClusters(profile model.ClusterProfile) ([]model.ClusterInfo, error) {
	config, err := clientcmd.LoadFromFile(profile.KubeConfigPath)
	if err != nil {
		return nil, err
	}

	var clusters []model.ClusterInfo

	for name, context := range config.Contexts {
		ci := model.ClusterInfo{
			Name:      name,
			Server:    context.Cluster,
			Current:   false,
			User:      context.AuthInfo,
			Namespace: context.Namespace,
		}
		clusters = append(clusters, ci)
	}

	return clusters, nil
}
