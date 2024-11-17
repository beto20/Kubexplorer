package endpoint

import (
	"Kubessistant/backend/objects"
	"Kubessistant/backend/usecase"
)

type IPodEndpoint interface {
	GetPods(namespace string) []objects.PodDto
}

type PodEndpoint struct {
	useCase usecase.IPodUseCase
}

func NewPodEndpoint(useCase usecase.IPodUseCase) IPodEndpoint {
	return &PodEndpoint{useCase: useCase}
}

func (pe *PodEndpoint) GetPods(namespace string) []objects.PodDto {
	return pe.useCase.GetAllPods(namespace)

	//for _, pod := range pods {
	//	fmt.Printf("pod name: %s namespace: %s requestCPU: %s limitsCPU: %s requestMemory: %s limitsMemory: %s status: %s age: %s\n",
	//		pod.Name,
	//		pod.Namespace,
	//		pod.Container.Request.Cpu,
	//		pod.Container.Limit.Cpu,
	//		pod.Container.Request.Memory,
	//		pod.Container.Limit.Memory,
	//		pod.Status,
	//		pod.Age,
	//	)
	//}
}
