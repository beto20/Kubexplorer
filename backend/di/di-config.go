package di

import (
	"Kubessistant/backend/container"
	"Kubessistant/backend/endpoint"
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

func SetupPodContainer() *container.Container {
	c := container.NewContainer()

	podObject := objects.NewPodObject()
	c.Register("IPod", podObject)

	podUseCase := usecase.NewPodUseCase(podObject)
	c.Register("IPod", podUseCase)

	podEndpoint := endpoint.NewPodEndpoint(podUseCase)
	c.Register("IPodEndpoint", podEndpoint)

	return c
}

func SetupDeploymentContainer() *container.Container {
	c := container.NewContainer()

	deploymentObject := objects.NewDeploymentObject()
	deploymentUseCase := usecase.NewDeploymentUseCase(deploymentObject)
	deploymentEndpoint := endpoint.NewDeploymentEndpoint(deploymentUseCase)
	c.Register("IDeploymentEndpoint", deploymentEndpoint)

	return c
}

func SetupServiceContainer() *container.Container {
	c := container.NewContainer()

	serviceObject := objects.NewServiceObject()
	serviceUseCase := usecase.NewServiceUseCase(serviceObject)
	serviceEndpoint := endpoint.NewServiceEndpoint(serviceUseCase)

	c.Register("IServiceEndpoint", serviceEndpoint)

	return c
}

func SetupEnvironmentContainer() *container.Container {
	c := container.NewContainer()

	environmentObject := objects.NewEnvironmentObject()
	environmentUseCase := usecase.NewEnvironmentUseCase(environmentObject)
	environmentEndpoint := endpoint.NewEnvironmentEndpoint(environmentUseCase)

	c.Register("IEnvironmentEndpoint", environmentEndpoint)

	return c
}
