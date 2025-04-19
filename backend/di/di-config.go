package di

import (
	"Kubessistant/backend/container"
	"Kubessistant/backend/database"
	"Kubessistant/backend/endpoint"
	"Kubessistant/backend/kubeclient"
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

func SetupWorkloadContainer() *container.Container {
	c := container.NewContainer()

	deploymentObject := objects.NewDeploymentObject()
	deploymentUseCase := usecase.NewDeploymentUseCase(deploymentObject)

	//podObject := objects.NewPodObject()
	podClient := kubeclient.NewPod()
	podUseCase := usecase.NewPodUseCase(podClient)

	workloadEndpoint := endpoint.NewWorkloadEndpoint(podUseCase, deploymentUseCase)

	c.Register("IWorkloadEndpoint", workloadEndpoint)

	return c
}

func SetupNetworkContainer() *container.Container {
	c := container.NewContainer()

	serviceObject := objects.NewServiceObject()
	serviceUseCase := usecase.NewServiceUseCase(serviceObject)

	ingressObject := objects.NewIngressObject()
	ingressUseCase := usecase.NewIngressUseCase(ingressObject)

	networkEndpoint := endpoint.NewNetworkEndpoint(serviceUseCase, ingressUseCase)

	c.Register("INetworkEndpoint", networkEndpoint)

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

func SetupParameterContainer() *container.Container {
	c := container.NewContainer()

	parameterEntity := database.NewParameterEntity()
	parameterUseCase := usecase.NewParameterUseCase(parameterEntity)
	parameterEndpoint := endpoint.NewParameterEndpoint(parameterUseCase)

	c.Register("IParameterEndpoint", parameterEndpoint)

	return c
}

func SetupGeneralContainer() *container.Container {
	c := container.NewContainer()

	nodeObject := objects.NewNodeObject()
	nodeUseCase := usecase.NewNodeUseCase(nodeObject)

	namespaceObject := objects.NewNamespaceObject()
	namespaceUseCase := usecase.NewNamespaceUseCase(namespaceObject)

	generalEndpoint := endpoint.NewGeneralEndpoint(nodeUseCase, namespaceUseCase)

	c.Register("IGeneralEndpoint", generalEndpoint)

	return c
}

func SetupStorageContainer() *container.Container {
	c := container.NewContainer()

	storageObject := objects.NewStorageObject()
	storageUseCase := usecase.NewStorageUseCase(storageObject)
	storageEndpoint := endpoint.NewStorageEndpoint(storageUseCase)

	c.Register("IStorageEndpoint", storageEndpoint)

	return c
}

//func SetupMetricsContainer() *container.Container {
//	c := container.NewContainer()
//
//	//podObject := objects.NewPodObject()
//	//nodeObject := objects.NewNodeObject()
//	metricService := service.NewMetricService()
//	metricBackground := background.NewMetricBackground(metricService)
//
//	c.Register("IMetricBackground", metricBackground)
//
//	return c
//}

func SetupMetricsContainer2() *container.Container {
	c := container.NewContainer()

	//podObject := objects.NewPodObject()
	//nodeObject := objects.NewNodeObject()
	metricUseCase := usecase.NewMetricUseCase()
	metricEndpoint := endpoint.NewMetricEndpoint(metricUseCase)

	c.Register("IMetricEndpoint", metricEndpoint)

	return c
}
