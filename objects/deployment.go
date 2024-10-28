package objects

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Deployment struct {
	Name      string
	Namespace string
	Status    string
	Age       string
}

func GetDeployments(namespace string) []Deployment {
	deploymentClient := client.AppsV1().Deployments(namespace)
	deploys, err := deploymentClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error when get deploys")
	}
	var deployments []Deployment

	for _, deploy := range deploys.Items {
		d := Deployment{
			Name:      deploy.Name,
			Namespace: deploy.Namespace,
			Status:    string(deploy.Status.Conditions[0].Status),
			Age:       deploy.CreationTimestamp.String(),
		}

		deployments = append(deployments, d)

		fmt.Printf("dep name: %s namespace: %s status: %s startTime: %s\n",
			deploy.Name,
			deploy.Namespace,
			deploy.Status.Conditions[0].Status,
			deploy.CreationTimestamp,
		)
	}

	return deployments
}
