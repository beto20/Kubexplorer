package endpoint

import (
	"Kubessistant/backend/database"
	"Kubessistant/backend/usecase"
)

type IParameterEndpoint interface {
	GetKubernetesParameters() []database.CommonParameterDto
	GetCommonParameters() []database.CommonParameterDto
	GetK8sObjects() []database.ObjectType
	GetHeadParams(k8sObject string) []database.HeadParamsDto
}

type parameterEndpoint struct {
	useCase usecase.IParameterUseCase
}

func NewParameterEndpoint(useCase usecase.IParameterUseCase) IParameterEndpoint {
	return &parameterEndpoint{useCase: useCase}
}

func (pe *parameterEndpoint) GetKubernetesParameters() []database.CommonParameterDto {
	return pe.useCase.GetKubernetesParameters()
}

func (pe *parameterEndpoint) GetCommonParameters() []database.CommonParameterDto {
	return pe.useCase.GetCommonParameters()
}

func (pe *parameterEndpoint) GetK8sObjects() []database.ObjectType {
	return pe.useCase.GetK8sObjects()
}

func (pe *parameterEndpoint) GetHeadParams(k8sObject string) []database.HeadParamsDto {
	if k8sObject == "pod" {
		return []database.HeadParamsDto{
			{Title: "Name", Key: "name", Align: "start", Sortable: true},
			{Title: "Namespace", Key: "namespace", Align: "start", Sortable: false},
			{Title: "Replicas", Key: "replicas", Align: "start", Sortable: false},
			{Title: "Cpu", Key: "cpu", Align: "start", Sortable: false},
			{Title: "Memory", Key: "memory", Align: "start", Sortable: false},
			{Title: "Age", Key: "age", Align: "start", Sortable: true},
			{Title: "Status", Key: "status", Align: "start", Sortable: false},
			{Title: "Actions", Key: "actions", Align: "start", Sortable: false},
		}
	} else if k8sObject == "node" {
		return []database.HeadParamsDto{
			{Title: "Name", Key: "name", Align: "start", Sortable: true},
			{Title: "Age", Key: "age", Align: "start", Sortable: true},
			{Title: "Cpu", Key: "cpu", Align: "start", Sortable: false},
			{Title: "Memory", Key: "memory", Align: "start", Sortable: false},
			{Title: "Status", Key: "status", Align: "start", Sortable: false},
			{Title: "Actions", Key: "actions", Align: "start", Sortable: false},
		}
	} else if k8sObject == "namespace" {
		return []database.HeadParamsDto{
			{Title: "Name", Key: "name", Align: "start", Sortable: true},
			//{Title: "Label", Key: "age", Align: "start", Sortable: true},
			{Title: "Age", Key: "age", Align: "start", Sortable: true},
			{Title: "Status", Key: "status", Align: "start", Sortable: false},
			{Title: "Actions", Key: "actions", Align: "start", Sortable: false},
		}
	} else if k8sObject == "deployment" {
		return []database.HeadParamsDto{
			{Title: "Name", Key: "name", Align: "start", Sortable: true},
			{Title: "Namespace", Key: "namespace", Align: "start", Sortable: false},
			{Title: "Replicas", Key: "replicas", Align: "start", Sortable: false},
			{Title: "Age", Key: "age", Align: "start", Sortable: true},
			{Title: "Status", Key: "status", Align: "start", Sortable: false},
			{Title: "Actions", Key: "actions", Align: "start", Sortable: false},
		}
	} else {
		return []database.HeadParamsDto{}
	}
}
