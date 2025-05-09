package kubeclient

import (
	"Kubessistant/backend/model"
	"context"
	"errors"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
)

type deploymentClient struct {
	client kubernetes.Interface
}

func NewDeployment(client kubernetes.Interface) DeploymentClient {
	return &deploymentClient{client: client}
}

func (d deploymentClient) GetDeployments() ([]model.DeploymentDto, error) {
	home, _ := os.UserHomeDir()
	kubeconfigPath := filepath.Join(home, ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		panic(err)
	}

	client := kubernetes.NewForConfigOrDie(config)

	deploymentClient := client.AppsV1().Deployments("TODO")
	deploys, err := deploymentClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error when get deploys")
	}
	var deployments []model.DeploymentDto

	for _, deploy := range deploys.Items {
		d := model.DeploymentDto{
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

	return deployments, errors.New("deployment list is empty")
}

func (d deploymentClient) GetDeploymentsMock() ([]model.DeploymentDto, error) {
	var deployments []model.DeploymentDto
	for i := 0; i < 10; i++ {

		p := model.DeploymentDto{
			Name:      fmt.Sprintf("Deployment %d", i),
			Namespace: "TODO",
			Status:    "Alive",
			Age:       strconv.Itoa(rand.Intn(1000)),
		}

		deployments = append(deployments, p)
	}

	return deployments, errors.New("Deployment Not Found")
}

func (d deploymentClient) GetDeployment(name string) (model.DeploymentDto, error) {
	//TODO implement me
	panic("implement me")
}

func (d deploymentClient) UpdateDeployment(name string) error {
	//TODO implement me
	panic("implement me")
}

func (d deploymentClient) DeleteDeployment(name string) error {
	//TODO implement me
	panic("implement me")
}
