package main

import (
	"Kubexplorer/backend/kubeclient"
	"fmt"
	"os"
	"path/filepath"
)

func getPath() string {
	home, _ := os.UserHomeDir()
	kubeConfigPath := filepath.Join(home, ".kube/config")

	//kubeConfig, _ := clientcmd.BuildConfigFromFlags("", kubeConfigPath)

	return kubeConfigPath
}

func main() {
	fmt.Println("Hello World")

	//client := kubeclient.NewPod(k8sClient)
	//client2 := kubeclient2.NewDeployment(k8sClient)
	//usecase := usecase2.NewPodUseCase(client)
	//usecase2 := usecase2.NewDeploymentUseCase(client2)
	//ep := endpoint.NewWorkloadEndpoint(usecase, usecase2)

	//ep.GetPod("")

	//x := kubeclient2.DeploymentClient{}
	//x.GetDeployments()

	//kubeConfig, _ := clientcmd.BuildConfigFromFlags("", getPath())

	conf := kubeclient.NewClusterManager()
	client, _ := conf.GetClient("minikube", getPath())

	y := kubeclient.NewPod(client)
	y.GetPods()
}
