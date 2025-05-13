package endpoint

import "Kubessistant/backend/usecase"

type MetricEndpoint struct {
	metricUseCase usecase.MetricUseCase
}

func NewMetricEndpoint(metricUseCase usecase.MetricUseCase) *MetricEndpoint {
	return &MetricEndpoint{metricUseCase: metricUseCase}
}

func (m *MetricEndpoint) GetPodMetric(namespace string) {
	//go async.PodMetricsThread(namespace)
	//return m.metricUseCase.GetPodMetric(namespace)
}
func (m *MetricEndpoint) GetNodeMetric(namespace string) {
	//go async.PodMetricsThread(namespace)
	//return m.metricUseCase.GetNodeMetric(namespace)
}
