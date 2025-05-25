package kubeclient

import (
	"Kubessistant/backend/model"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type cluster struct {
	client kubernetes.Interface
}

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

func (c *cluster) GetCurrentCluster() (model.EnvironmentDto, error) {
	dto := model.EnvironmentDto{
		Name:        "minikube",
		Description: "minikube description",
		Env:         "Dev",
		Status:      true,
	}

	return dto, nil
}

func (c *cluster) GetClusters() ([]model.EnvironmentDto, error) {
	return []model.EnvironmentDto{
		{
			Name:        "minikube",
			Description: "minikube description",
			Env:         "Dev",
			Status:      true,
		},
		{
			Name:        "minikube-2",
			Description: "minikube description",
			Env:         "UAT",
			Status:      true,
		},
		{
			Name:        "minikube-3",
			Description: "minikube description",
			Env:         "PRD",
			Status:      false,
		},
		{
			Name:        "minikube-4",
			Description: "minikube description",
			Env:         "PRD-2",
			Status:      true,
		},
		{
			Name:        "minikube-5",
			Description: "minikube description",
			Env:         "Sandbox",
			Status:      false,
		},
	}, nil
}

func (c *cluster) GetNode(name string) (model.NodeDtoV2, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	node, err := clientset.CoreV1().Nodes().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	return model.NodeDtoV2{
		Name:      node.Name,
		Namespace: node.Namespace,
		Resource: model.Resource{
			Cpu:    node.Status.Capacity.Cpu().String(),
			Memory: node.Status.Capacity.Memory().String(),
		},
		Version:           node.ResourceVersion,
		CreationTimestamp: node.CreationTimestamp.String(),
		Labels:            node.Labels,
	}, nil
}

func (c *cluster) GetNodes() ([]model.NodeDto, error) {
	return []model.NodeDto{
		{
			Name: "node-1",
			Resource: model.Resource{
				Cpu:    "6Gi",
				Memory: "20Gi",
			},
			Roles: []string{
				"ADMIN",
				"GENERAL",
			},
			Version: "1.24.0",
			Age:     "20 day",
			Status:  true,
		},
		{
			Name: "node-2",
			Resource: model.Resource{
				Cpu:    "6Gi",
				Memory: "20Gi",
			},
			Roles: []string{
				"ADMIN",
				"GENERAL",
			},
			Version: "1.28.0",
			Age:     "100 day",
			Status:  true,
		},
	}, nil
}
