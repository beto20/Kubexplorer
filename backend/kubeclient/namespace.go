package kubeclient

import (
	"Kubessistant/backend/model"
	"k8s.io/client-go/kubernetes"
)

type namespaceClient struct {
	client kubernetes.Interface
}

func NewNamespaceClient(client kubernetes.Interface) NamespaceClient {
	return &namespaceClient{client: client}
}

func (n namespaceClient) GetNamespaces() ([]model.NamespaceDto, error) {
	return []model.NamespaceDto{
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

func (n namespaceClient) GetNamespace(name string) (model.NamespaceDto, error) {
	return model.NamespaceDto{
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
	}, nil
}

func (n namespaceClient) UpdateNamespace(name string) error {
	//TODO implement me
	panic("implement me")
}

func (n namespaceClient) DeleteNamespace(name string) error {
	//TODO implement me
	panic("implement me")
}
