package metrics

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
)

type MetricWrapper interface {
	PushToSummarytMetrics() func(*prometheus.SummaryVec, string, *error, context.Context)
	PushToErrorCounterMetrics() func(*prometheus.CounterVec, error, context.Context)
}
