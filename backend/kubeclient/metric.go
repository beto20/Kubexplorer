package kubeclient

import (
	"Kubessistant/backend/model"
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

type metric struct {
	client metricsv.Interface
}

func NewMetric() MetricClient {
	return &metric{}
}

func (m *metric) GetPodMetrics(namespace string, chMetricDto <-chan []model.PodMetricDto) []model.PodMetricDto {
	//inClusterConfig, err := rest.InClusterConfig()
	//if err != nil {
	//	fmt.Printf("Error creating in-cluster config: %v", err)
	//}

	//metricsClient, err := metricsv.NewForConfig(inClusterConfig)
	//if err != nil {
	//	fmt.Print(err)
	//}

	return pollMetrics(m.client, namespace)
}

func pollMetrics(metricsClient metricsv.Interface, namespace string) []model.PodMetricDto {
	podMetrics, err := metricsClient.MetricsV1beta1().PodMetricses(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error to get pod metrics")
	}

	var metrics []model.PodMetricDto

	for _, podMetric := range podMetrics.Items {
		for _, container := range podMetric.Containers {
			metrics = append(metrics, model.PodMetricDto{
				Name:      podMetric.Name,
				Namespace: podMetric.Namespace,
				Consume: model.Resource{
					Cpu:              container.Usage.Cpu().String(),
					Memory:           container.Usage.Memory().String(),
					Storage:          container.Usage.Storage().String(),
					StorageEphemeral: container.Usage.StorageEphemeral().String(),
				},
			})
		}
	}

	return metrics
}
