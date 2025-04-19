package middleware

import (
	"Kubessistant/backend/database"
	"Kubessistant/backend/endpoint"
)

type ParameterMiddleware struct {
	endpoint endpoint.IParameterEndpoint
}

func NewParameterMiddleware(endpoint endpoint.IParameterEndpoint) *ParameterMiddleware {
	return &ParameterMiddleware{endpoint}
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
