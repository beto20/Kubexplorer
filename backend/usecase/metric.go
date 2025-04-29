package usecase

type IMetricUseCase interface {
	GetPodMetric(namespace string)
	GetNodeMetric(namespace string)
}

type metricImpl struct {
}

func NewMetricUseCase() IMetricUseCase {
	return &metricImpl{}
}

func (m *metricImpl) GetPodMetric(namespace string) {
	//return m.service.GetPodMetrics(namespace)
}

func (m *metricImpl) GetNodeMetric(namespace string) {
	//return m.service.GetNodeMetrics(namespace)
}
