package objects

import (
	"fmt"
	"math/rand"
	"strconv"
)

type IDeploymentObject interface {
	GetDeployments() []DeploymentDto
	GetDeploymentsMock() []DeploymentDto
	GetDeploymentByName(name string) DeploymentDto
	UpdateDeploymentByName(name string) bool
	DeleteDeploymentByName(name string) bool
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

func (d *deploymentImpl) GetDeployments() []DeploymentDto {
	//deploymentClient := client.AppsV1().Deployments("TODO")
	//deploys, err := deploymentClient.List(context.TODO(), metav1.ListOptions{})
	//if err != nil {
	//	fmt.Println("Error when get deploys")
	//}
	var deployments []DeploymentDto

	//for _, deploy := range deploys.Items {
	//	d := DeploymentDto{
	//		Name:      deploy.Name,
	//		Namespace: deploy.Namespace,
	//		Status:    string(deploy.Status.Conditions[0].Status),
	//		Age:       deploy.CreationTimestamp.String(),
	//	}
	//
	//	deployments = append(deployments, d)
	//
	//	fmt.Printf("dep name: %s namespace: %s status: %s startTime: %s\n",
	//		deploy.Name,
	//		deploy.Namespace,
	//		deploy.Status.Conditions[0].Status,
	//		deploy.CreationTimestamp,
	//	)
	//}

	return deployments
}

func (d *deploymentImpl) GetDeploymentsMock() []DeploymentDto {
	var deployments []DeploymentDto
	for i := 0; i < 10; i++ {

		p := DeploymentDto{
			Name:      fmt.Sprintf("Deployment %d", i),
			Namespace: "TODO",
			Status:    "Alive",
			Age:       strconv.Itoa(rand.Intn(1000)),
		}

		deployments = append(deployments, p)
	}

	return deployments
}

func (d *deploymentImpl) GetDeploymentByName(name string) DeploymentDto {
	var deployments DeploymentDto

	return deployments
}

func (d *deploymentImpl) UpdateDeploymentByName(name string) bool {
	return true
}

func (d *deploymentImpl) DeleteDeploymentByName(name string) bool {
	return true
}
