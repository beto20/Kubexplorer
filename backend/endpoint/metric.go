package endpoint

import "Kubessistant/backend/usecase"

type IMetricEndpoint interface {
	GetPodMetric(namespace string)
	GetNodeMetric(namespace string)
}

type metricEndpoint struct {
	metricUseCase usecase.IMetricUseCase
}

func NewMetricEndpoint(metricUseCase usecase.IMetricUseCase) IMetricEndpoint {
	return &metricEndpoint{metricUseCase: metricUseCase}
}

func (m *metricEndpoint) GetPodMetric(namespace string) {
	//go async.PodMetricsThread(namespace)
	//return m.metricUseCase.GetPodMetric(namespace)
}
func (m *metricEndpoint) GetNodeMetric(namespace string) {
	//go async.PodMetricsThread(namespace)
	//return m.metricUseCase.GetNodeMetric(namespace)
}
