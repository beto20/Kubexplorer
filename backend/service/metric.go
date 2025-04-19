package service

import (
	"context"
	"fmt"
	"time"
)

type IMetricService interface {
	CollectPodMetrics(ctx context.Context, namespace string, interval time.Duration)
	CollectNodeMetrics()
	GetPodMetrics()
	GetNodeMetrics()
}

type metricImpl struct {
}

func NewMetricService() IMetricService {
	return &metricImpl{}
}

// func (m *metricImpl) CollectPodMetrics(ctx context.Context, namespace string, out chan<- []objects.PodMetricDto, interval time.Duration) {
func (m *metricImpl) CollectPodMetrics(ctx context.Context, namespace string, interval time.Duration) {
	fmt.Println("Pod metrics background process started")
	//metricsChan := make(chan []objects.PodMetricDto)

	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	/////
	//timer := time.NewTicker(interval)
	//defer timer.Stop()
	//
	//for {
	//	select {
	//	case <-ctx.Done():
	//		fmt.Println("Stopping polling goroutine")
	//		return
	//	case <-timer.C:
	//		metrics := m.pod.GetPodMetric(namespace)
	//		out <- metrics
	//	}
	//}
}

func (m *metricImpl) CollectNodeMetrics() {
}

func (m *metricImpl) GetPodMetrics() {
}

func (m *metricImpl) GetNodeMetrics() {
}
