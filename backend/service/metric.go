package service

type IMetricService interface {
	GetPodMetric()
	GetNodeMetric()
}

type metricImpl struct {
}

func NewMetricService() IMetricService {
	return &metricImpl{}
}

func (m *metricImpl) GetPodMetric()  {}
func (m *metricImpl) GetNodeMetric() {}
