package middleware

import "Kubessistant/backend/endpoint"

type MetricMiddleware struct {
	endpoint endpoint.IMetricEndpoint
}

func NewMetricMiddleware(endpoint endpoint.IMetricEndpoint) *MetricMiddleware {
	return &MetricMiddleware{endpoint: endpoint}
}
