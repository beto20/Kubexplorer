package main

import (
	"Kubessistant/backend/kubeclient2"
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	//client := kubeclient.NewPod(k8sClient)
	//client2 := kubeclient2.NewDeployment(k8sClient)
	//usecase := usecase2.NewPodUseCase(client)
	//usecase2 := usecase2.NewDeploymentUseCase(client2)
	//ep := endpoint.NewWorkloadEndpoint(usecase, usecase2)

	//ep.GetPod("")

	x := kubeclient2.DeploymentClient{}
	x.GetDeployments()
}
