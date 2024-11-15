package di

import (
	"github.com/beto20/kubessitant/container"
	"github.com/beto20/kubessitant/endpoint"
	"github.com/beto20/kubessitant/objects"
	"github.com/beto20/kubessitant/usecase"
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
