package kubeclient

import (
	"Kubessistant/backend/model"
)

type metric struct{}

func NewMetric() MetricClient {
	return &metric{}
}

func (m *metric) GetPodMetrics(namespace string) []model.PodMetricDto {
	//inClusterConfig, err := rest.InClusterConfig()
	//if err != nil {
	//	fmt.Printf("Error creating in-cluster config: %v", err)
	//}
	//
	//metricsClient, err := metricsv.NewForConfig(inClusterConfig)
	//if err != nil {
	//	fmt.Print(err)
	//}
	//
	//return pollMetrics(metricsClient, namespace)

	var metrics []model.PodMetricDto
	return metrics
}

//func pollMetrics(metricsClient *metricsv.Clientset, namespace string) []model.PodMetricDto {
//podMetrics, err := metricsClient.MetricsV1beta1().PodMetricses(namespace).List(context.TODO(), metav1.ListOptions{})
//if err != nil {
//	fmt.Println("Error to get pod metrics")
//}

//var metrics []model.PodMetricDto

//for _, podMetric := range podMetrics.Items {
//	for _, container := range podMetric.Containers {
//		metrics = append(metrics, model.PodMetricDto{
//			Name:      podMetric.Name,
//			Namespace: podMetric.Namespace,
//			Consume: model.Resource{
//				Cpu:              container.Usage.Cpu().String(),
//				Memory:           container.Usage.Memory().String(),
//				Storage:          container.Usage.Storage().String(),
//				StorageEphemeral: container.Usage.StorageEphemeral().String(),
//			},
//		})
//	}
//}

//return metrics
//}
