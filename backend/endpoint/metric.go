package endpoint

import "Kubessistant/backend/async"

type IMetricEndpoint interface {
	GetPodMetric(namespace string)
	GetNodeMetric(namespace string)
}

type metricEndpoint struct {
}

func NewMetricEndpoint() IMetricEndpoint {
	return &metricEndpoint{}
}

func (m *metricEndpoint) GetPodMetric(namespace string) {
	go async.PodMetricsThread(namespace)
}
func (m *metricEndpoint) GetNodeMetric(namespace string) {
	go async.PodMetricsThread(namespace)
}
