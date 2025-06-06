package service

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
	"context"
	"fmt"
	"time"
)

type MetricService interface {
	CollectPodMetrics(ctx context.Context, namespace string, out chan<- []model.PodMetricDto, interval time.Duration)
	CollectNodeMetrics()
	GetPodMetrics()
	GetNodeMetrics()
}

type metricService struct {
	metric kubeclient.MetricClient
}

func NewMetricService(metric kubeclient.MetricClient) MetricService {
	return &metricService{metric: metric}
}

func Start() {
	x := NewMetricService(nil)
	x.CollectPodMetrics(context.Background(), "", make(chan<- []model.PodMetricDto), time.Second*5)

}

func (m *metricService) CollectPodMetrics(ctx context.Context, namespace string, out chan<- []model.PodMetricDto, interval time.Duration) {
	fmt.Println("Pod metrics background process started")
	metricsChan := make(chan []model.PodMetricDto)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timer := time.NewTicker(interval)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopping polling goroutine")
			return
		case <-timer.C:
			metrics := m.metric.GetPodMetrics(namespace, metricsChan)
			out <- metrics
		}
	}
}

func (m *metricService) CollectNodeMetrics() {
}

func (m *metricService) GetPodMetrics() {
}

func (m *metricService) GetNodeMetrics() {
}
