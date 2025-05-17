package middleware

import (
	"Kubessistant/backend/database"
	"Kubessistant/backend/endpoint"
	"Kubessistant/backend/usecase"
	"k8s.io/client-go/kubernetes"
)

type ParameterMiddleware struct {
	endpoint endpoint.ParameterEndpoint
}

func NewParameterMiddleware(endpoint *endpoint.ParameterEndpoint) *ParameterMiddleware {
	return &ParameterMiddleware{endpoint: *endpoint}
}

func (p *ParameterMiddleware) GetKubernetesParameters() []database.CommonParameterDto {
	return p.endpoint.GetKubernetesParameters()
}

func (p *ParameterMiddleware) GetCommonParameters() []database.CommonParameterDto {
	return p.endpoint.GetCommonParameters()
}

func (p *ParameterMiddleware) GetK8sObjects() []database.ObjectType {
	return p.endpoint.GetK8sObjects()
}

func (p *ParameterMiddleware) GetHeaderParams(k8sObject string) []database.HeadParamsDto {
	return p.endpoint.GetHeadParams(k8sObject)
}

func BuildParameters(client kubernetes.Interface) *ParameterMiddleware {
	paramDatabase := database.NewParameterEntity()
	parameterUseCase := usecase.NewParameterUseCase(paramDatabase)

	parameterEndpoint := endpoint.NewParameterEndpoint(parameterUseCase)

	return NewParameterMiddleware(parameterEndpoint)
}
