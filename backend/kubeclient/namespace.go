package kubeclient

import (
	"Kubexplorer/backend/model"
	"context"
	"errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type namespaceClient struct {
	client kubernetes.Interface
}

func NewNamespaceClient(client kubernetes.Interface) NamespaceClient {
	return &namespaceClient{client: client}
}

func (n namespaceClient) GetNamespaces() ([]model.NamespaceDto, error) {
	namespaces, err := n.client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic("Failed to list namespaces")
	}

	var result []model.NamespaceDto

	for _, ns := range namespaces.Items {
		dto := model.NamespaceDto{
			Name:         ns.Name,
			Version:      ns.APIVersion,
			CreationTime: ns.CreationTimestamp.String(),
			Labels:       ns.Labels,
			Status:       ns.Status.String(),
		}
		result = append(result, dto)
	}

	return result, errors.New("No namespaces found")
}

func (n namespaceClient) GetNamespace(name string) (model.NamespaceDto, error) {
	ns, err := n.client.CoreV1().Namespaces().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic("Failed to list namespaces")
	}

	return model.NamespaceDto{
		Name:         ns.Name,
		Version:      ns.APIVersion,
		CreationTime: ns.CreationTimestamp.String(),
		Labels:       ns.Labels,
		Status:       ns.Status.String(),
	}, errors.New("No namespace found")
}

func (n namespaceClient) UpdateNamespace(name string, dto model.NamespaceDto) error {
	client := n.client.CoreV1().Namespaces()
	ns, err := client.Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic("Error while searching ingress")
	}

	ns.Name = dto.Name

	_, err = client.Update(context.TODO(), ns, metav1.UpdateOptions{})

	return err
}

func (n namespaceClient) DeleteNamespace(name string) error {
	return n.client.CoreV1().Namespaces().Delete(context.TODO(), name, metav1.DeleteOptions{})
}
