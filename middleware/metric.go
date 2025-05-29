package middleware

import "Kubexplorer/backend/endpoint"

type MetricMiddleware struct {
	endpoint endpoint.MetricEndpoint
}

func NewMetricMiddleware(endpoint endpoint.MetricEndpoint) *MetricMiddleware {
	return &MetricMiddleware{endpoint: endpoint}
}
