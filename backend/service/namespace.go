package service

import (
	"context"
	"fmt"
	"gopkg.in/yaml.v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"os"
)

type NamespaceService interface {
	ExportObjects(namespace string, directory string) error
}

type namespaceService struct {
	dynamicClient dynamic.Interface
}

func NewNamespaceService(dynamicClient dynamic.Interface) NamespaceService {
	return &namespaceService{dynamicClient: dynamicClient}
}

func (ns *namespaceService) ExportObjects(namespace string, directory string) error {

	resourceTypes := []schema.GroupVersionResource{
		{Group: "apps", Version: "v1", Resource: "deployments"},
		{Group: "", Version: "v1", Resource: "services"},
		{Group: "", Version: "v1", Resource: "configmaps"},
		{Group: "", Version: "v1", Resource: "secrets"},
		{Group: "", Version: "v1", Resource: "pods"},
	}

	for _, r := range resourceTypes {
		resList, _ := ns.dynamicClient.Resource(r).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
		for _, item := range resList.Items {
			yamlBytes, _ := yaml.Marshal(item.Object)
			filePath := fmt.Sprintf("snapshots/%s/%s-%s.yaml", namespace, r.Resource, item.GetName())
			os.WriteFile(filePath, yamlBytes, 0644)
		}
	}
	return nil
}
