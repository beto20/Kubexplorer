package usecase

type IMetricUseCase interface {
	GetPodMetric()
	GetNodeMetric()
}

type metricImpl struct {
}

func NewMetricUseCase() IMetricUseCase {
	return &metricImpl{}
}

func (m *metricImpl) GetPodMetric() {

}

func (m *metricImpl) GetNodeMetric() {

}
