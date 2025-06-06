package background

//type IMetricBackground interface {
//}
//
//type metricImpl struct {
//	service service.IMetricService
//}
//
//func NewMetricBackground(service service.IMetricService) IMetricBackground {
//	return &metricImpl{service: service}
//}

//func (m MetricBackground) PodMetricsBackgroundProcess() {
//	fmt.Println("Pod metrics background process started")
//	metricsChan := make(chan []objects.PodMetricDto)
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	go m.service.CollectPodMetrics(ctx, "default", metricsChan, 10*time.Second)
//
//}
