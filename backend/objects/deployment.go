package objects

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"math/rand"
	"strconv"
)

type IDeploymentObject interface {
	GetDeployments(namespace string) []DeploymentDto
	GetDeploymentsMock(namespace string) []DeploymentDto
}

type deploymentImpl struct{}

func NewDeploymentObject() IDeploymentObject {
	return &deploymentImpl{}
}

type DeploymentDto struct {
	Name      string
	Namespace string
	Status    string
	Age       string
}

func (d *deploymentImpl) GetDeployments(namespace string) []DeploymentDto {
	deploymentClient := client.AppsV1().Deployments(namespace)
	deploys, err := deploymentClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error when get deploys")
	}
	var deployments []DeploymentDto

	for _, deploy := range deploys.Items {
		d := DeploymentDto{
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

func (d *deploymentImpl) GetDeploymentsMock(namespace string) []DeploymentDto {
	var deployments []DeploymentDto
	for i := 0; i < 10; i++ {

		p := DeploymentDto{
			Name:      fmt.Sprintf("Deployment %d", i),
			Namespace: namespace,
			Status:    "Alive",
			Age:       strconv.Itoa(rand.Intn(1000)),
		}

		deployments = append(deployments, p)
	}

	return deployments
}
