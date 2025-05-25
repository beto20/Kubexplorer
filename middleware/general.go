package middleware

import (
	"Kubessistant/backend/endpoint"
	"Kubessistant/backend/kubeclient"
	"Kubessistant/backend/model"
	"Kubessistant/backend/usecase"
	"k8s.io/client-go/kubernetes"
)

type GeneralMiddleware struct {
	endpoint endpoint.GeneralEndpoint
}

func NewGeneralMiddleware(endpoint *endpoint.GeneralEndpoint) *GeneralMiddleware {
	return &GeneralMiddleware{endpoint: *endpoint}
}

func (g *GeneralMiddleware) GetNodes() ([]model.NodeDto, error) {
	return g.endpoint.GetNodes()
}

func (g *GeneralMiddleware) GetNodeByName(name string) (model.NodeDtoV2, error) {
	return g.endpoint.GetNodeByName(name)
}

func (g *GeneralMiddleware) GetNamespace() ([]model.NamespaceDto, error) {
	return g.endpoint.GetNamespace()
}

func (g *GeneralMiddleware) GetNamespaceByName(name string) (model.NamespaceDto, error) {
	return g.endpoint.GetNamespaceByName(name)
}

func (g *GeneralMiddleware) UpdateNamespaceByName(name string, dto model.NamespaceDto) error {
	return g.endpoint.UpdateNamespaceByName(name, dto)
}

func (g *GeneralMiddleware) DeleteNamespaceByName(name string) error {
	return g.endpoint.DeleteNamespaceByName(name)
}

func BuildGeneral(client kubernetes.Interface) *GeneralMiddleware {
	nodeClient := kubeclient.NewCluster()
	namespaceClient := kubeclient.NewNamespaceClient(client)

	nodeUseCase := usecase.NewNodeUseCase(nodeClient)
	namespaceUseCase := usecase.NewNamespaceUseCase(namespaceClient)

	generalEndpoint := endpoint.NewGeneralEndpoint(nodeUseCase, namespaceUseCase)

	return NewGeneralMiddleware(generalEndpoint)
}
