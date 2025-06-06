package kubeclient

import (
	"Kubexplorer/backend/model"
	"context"
	"errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ingressClient struct {
	client kubernetes.Interface
}

func NewIngressClient(client kubernetes.Interface) IngressClient {
	return &ingressClient{client: client}
}

func (i ingressClient) GetIngresses() ([]model.IngressDto, error) {
	ingresses, err := i.client.NetworkingV1().Ingresses("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic("Error while listing ingresses")
	}

	var result []model.IngressDto

	for _, ingress := range ingresses.Items {
		var rulesDto []model.RuleDto

		for _, rule := range ingress.Spec.Rules {
			dto := model.RuleDto{
				Host: rule.Host,
				Path: rule.HTTP.Paths[0].Path,
			}
			rulesDto = append(rulesDto, dto)
		}

		dto := model.IngressDto{
			Name:      ingress.GetName(),
			Namespace: ingress.GetNamespace(),
			Creation:  ingress.GetCreationTimestamp().String(),
			Labels:    ingress.GetLabels(),
			Rules:     rulesDto,
		}
		result = append(result, dto)
	}

	return result, errors.New("Error while listing ingresses")
}

func (i ingressClient) GetIngress(name string, namespace string) (model.IngressDto, error) {
	ingress, _ := i.client.NetworkingV1().Ingresses(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	var rulesDto []model.RuleDto

	for _, rule := range ingress.Spec.Rules {
		dto := model.RuleDto{
			Host: rule.Host,
			Path: rule.HTTP.Paths[0].Path,
		}
		rulesDto = append(rulesDto, dto)
	}

	return model.IngressDto{
		Name:      ingress.GetName(),
		Namespace: ingress.GetNamespace(),
		Creation:  ingress.GetCreationTimestamp().String(),
		Labels:    ingress.GetLabels(),
		Rules:     rulesDto,
	}, errors.New("Error while listing ingresses")
}

func (i ingressClient) UpdateIngress(name string, namespace string, dto model.IngressDto) error {
	client := i.client.NetworkingV1().Ingresses(namespace)
	ingress, err := client.Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic("Error while searching ingress")
	}

	ingress.Name = dto.Name
	ingress.Namespace = dto.Namespace
	ingress.Labels = dto.Labels

	_, err = client.Update(context.TODO(), ingress, metav1.UpdateOptions{})

	return err
}

func (i ingressClient) DeleteIngress(name string, namespace string) error {
	return i.client.NetworkingV1().Ingresses(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}
