package kubeclient

import (
	"Kubexplorer/backend/model"
	"context"
	"errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type serviceClient struct {
	client kubernetes.Interface
}

func NewServiceClient(client kubernetes.Interface) ServiceClient {
	return &serviceClient{client: client}
}

func (s serviceClient) GetServices(namespace string) ([]model.ServiceDto, error) {
	services, err := s.client.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic("Failed to list services")
	}

	var result []model.ServiceDto

	for _, service := range services.Items {
		dto := model.ServiceDto{
			Name:              service.Name,
			Namespace:         service.Namespace,
			Labels:            service.Labels,
			Status:            service.Status.String(),
			CreationTimestamp: service.CreationTimestamp.String(),
			Spec:              service.Spec.String(),
		}

		result = append(result, dto)
	}

	return result, errors.New("No services found")
}

func (s serviceClient) GetService(name string, namespace string) (model.ServiceDto, error) {
	service, err := s.client.CoreV1().Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic("Failed to get service")
	}

	return model.ServiceDto{
		Name:              service.Name,
		Namespace:         service.Namespace,
		Labels:            service.Labels,
		Status:            service.Status.String(),
		CreationTimestamp: service.CreationTimestamp.String(),
		Spec:              service.Spec.String(),
	}, errors.New("No service found")
}

func (s serviceClient) UpdateService(name string, namespace string, dto model.ServiceDto) error {
	client := s.client.CoreV1().Services(namespace)
	service, err := client.Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic("Error while searching ingress")
	}

	service.Name = dto.Name
	service.Namespace = dto.Namespace
	service.Labels = dto.Labels

	_, err = client.Update(context.TODO(), service, metav1.UpdateOptions{})

	return err
}

func (s serviceClient) DeleteService(name string, namespace string) error {
	return s.client.CoreV1().Services(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}
